package middleware

import (
    "log"
    "net/http"
)

// LogRequest is a middleware handler that logs the request path
func LogRequest(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        log.Printf("Received request for %s", r.URL.Path)
        next.ServeHTTP(w, r)
    })
}

// ErrorHandler is a middleware to catch and log errors
func ErrorHandler(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        defer func() {
            if err := recover(); err != nil {
                log.Printf("Recovered from an error: %s", err)
                http.Error(w, "Internal Server Error", http.StatusInternalServerError)
            }
        }()
        next.ServeHTTP(w, r)
    })
}
