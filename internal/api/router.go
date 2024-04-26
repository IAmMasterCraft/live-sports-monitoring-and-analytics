package api

import (
    "net/http"
)

// RegisterRoutes sets up the routing for the API endpoints
func RegisterRoutes() *http.ServeMux {
    mux := http.NewServeMux()
    mux.HandleFunc("/", homeHandler)
    // Add more handlers here
    return mux
}

// homeHandler handles the root path
func homeHandler(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/" {
        http.NotFound(w, r)
        return
    }
    w.Write([]byte("Welcome to the Live Event Monitoring Dashboard!"))
}
