package helpers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Sirupsen/logrus"
)

type responseWriterWrapper struct {
	impl       http.ResponseWriter
	statusCode int
}

func (p *responseWriterWrapper) Header() http.Header {
	return p.impl.Header()
}

func (p *responseWriterWrapper) Write(bytes []byte) (int, error) {
	return p.impl.Write(bytes)
}

func (p *responseWriterWrapper) WriteHeader(statusCode int) {
	p.statusCode = statusCode
	p.impl.WriteHeader(statusCode)
}

type loggingHandler struct {
	next http.Handler
}

func (p *loggingHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()

	wrapper := &responseWriterWrapper{w, -1}

	p.next.ServeHTTP(wrapper, r)

	statusCode := wrapper.statusCode
	responseTime := int64(time.Since(startTime) / time.Millisecond)

	logrus.WithFields(logrus.Fields{
		"request_time": startTime.Format(time.RFC3339),
		"remote_addr":  r.RemoteAddr,
		"method":       r.Method,
		"path":         r.URL.Path,
		"status":       statusCode,
		"time_taken":   fmt.Sprintf("%dms", responseTime),
	}).Info(http.StatusText(statusCode))
}

// NewHandler returns a middleware which logs requests and responses.
func NewHandler(next http.Handler) http.Handler {
	return &loggingHandler{next}
}
