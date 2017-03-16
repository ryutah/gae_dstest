package middleware

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func NewRequestLogger() RequestLogger {
	return new(requestLogger)
}

type RequestLogger interface {
	ServeHTTP(http.ResponseWriter, *http.Request, http.HandlerFunc)
}

type requestLogger struct {
}

func (rh *requestLogger) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	next(w, r)
}

func (rh *requestLogger) printHeaders(r *http.Request) {
	var headerStrs []string
	for header, values := range r.Header {
		s := fmt.Sprintf("%v: %v", header, strings.Join(values, ","))
		headerStrs = append(headerStrs, s)
	}
	log.Println("[info] %v", strings.Join(headerStrs, "; "))
}
