package api

import (
    "fmt"
    "net/http"
)

// HandleRequests is the entry point for incoming web requests to the API.
func HandleRequests(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Welcome to the Live Event Monitoring Dashboard!")
}
