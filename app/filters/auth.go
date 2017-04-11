package filters

import "net/http"

type basicAuth struct {
	next http.Handler
}

func (b *basicAuth) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	username, password, ok := r.BasicAuth()
	if ok == true && username == "wise" && password == "fuckfuckfuck" {
		b.next.ServeHTTP(w, r)
		return
	}
	w.Header().Set("WWW-Authenticate", `Basic realm="Wise Auth"`)
	w.WriteHeader(401)
	w.Write([]byte("401 Unauthorized\n"))
}

// BasicAuthHandler ...
func BasicAuthHandler(next http.Handler) http.Handler {
	return &basicAuth{next}
}
