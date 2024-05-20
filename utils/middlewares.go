package utils

import (
	"log"
	"net/http"
	"time"
)

// Override the HTTP method in case the client does not support the method
func MethodOverride(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			log.Printf("Current method is %s, override to %s\n", r.Method, r.FormValue("_method"))
			if method := r.FormValue("_method"); method != "" {
				r.Method = method
			}
			log.Printf("New method is %s\n", r.Method)
		}
		next.ServeHTTP(w, r)
	})
}

// LoggingMiddleware logs the details of each request
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("%s %s %s %s", r.Method, r.RequestURI, r.RemoteAddr, time.Since(start))
		next.ServeHTTP(w, r)
	})
}
