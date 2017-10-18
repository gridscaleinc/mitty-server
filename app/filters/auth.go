package filters

import (
	"context"
	"net/http"

	"mitty.co/mitty-server/app/models"
)

const (
	apiAuthContextKey = "filters/api_auth"
)

type basicAuth struct {
	next http.Handler
}

type apiAuth struct {
	next http.Handler
}

type apiKey struct {
	next http.Handler
}

func (b *basicAuth) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	username, password, ok := r.BasicAuth()
	if ok == true && username == "mitty" && password == "mittymitty" {
		b.next.ServeHTTP(w, r)
		return
	}
	w.Header().Set("WWW-Authenticate", `Basic realm="Wise Auth"`)
	w.WriteHeader(401)
	w.Write([]byte("401 Unauthorized\n"))
}

func (a *apiAuth) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	apiKey := r.Header.Get("X-Mitty-APIKEY")
	if apiKey != "pdXQWU2EpNMFPoCr6UAdMNUevAzuuG" {
		w.WriteHeader(403)
		w.Write([]byte("403 Forbidden\n"))
		return
	}
	accessToken := r.Header.Get("X-Mitty-AccessToken")
	user, err := models.GetUserByAccessToken(accessToken)
	if err != nil || user == nil {
		w.WriteHeader(401)
		w.Write([]byte("401 Unauthorized\n"))
		return
	}
	r = r.WithContext(context.WithValue(r.Context(), apiAuthContextKey, user.ID))
	a.next.ServeHTTP(w, r)
	return
}

func (a *apiKey) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	apiKey := r.Header.Get("X-Mitty-APIKEY")
	if apiKey != "pdXQWU2EpNMFPoCr6UAdMNUevAzuuG" {
		w.WriteHeader(403)
		w.Write([]byte("403 Forbidden\n"))
		return
	}
	a.next.ServeHTTP(w, r)
	return
}

// BasicAuthHandler ...
func BasicAuthHandler(next http.Handler) http.Handler {
	return &basicAuth{next}
}

// APIAuthHandler ...
func APIAuthHandler(next http.Handler) http.Handler {
	return &apiAuth{next}
}

// APIKeyHandler ...
func APIKeyHandler(next http.Handler) http.Handler {
	return &apiKey{next}
}

// GetCurrentUserID ...
//  Get the authorized user id.
//   return 0 if thre request context has not an authorized user.
func GetCurrentUserID(r *http.Request) int {
	value := r.Context().Value(apiAuthContextKey)
	if value == nil {
		return 0
	}
	return value.(int)
}
