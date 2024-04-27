package api

import (
    "net/http"
	"live-event-dashboard/internal/store"
)

// RegisterRoutes sets up the routing for the API endpoints
func RegisterRoutes(db *store.Database) *http.ServeMux {
    mux := http.NewServeMux()
    mux.HandleFunc("/", homeHandler)
    // Add more handlers here
	mux.HandleFunc("/events", GetLiveEvents(db))
    mux.HandleFunc("/events/details", GetEventDetails(db))
    mux.HandleFunc("/events/update", UpdateEvent(db))
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
