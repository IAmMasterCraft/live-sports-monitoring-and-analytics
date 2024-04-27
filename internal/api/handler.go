package api

import (
    "fmt"
    "net/http"
    "encoding/json"
    "live-event-dashboard/internal/model"
    "live-event-dashboard/internal/store"
    "strconv"
)

// HandleRequests is the entry point for incoming web requests to the API.
func HandleRequests(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Welcome to the Live Event Monitoring Dashboard!")
}

// Get live events
func GetLiveEvents(db *store.Database) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        events, err := db.GetLiveEvents()
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        json.NewEncoder(w).Encode(events)
    }
}

// Get event details
func GetEventDetails(db *store.Database) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Example: Extracting ID from URL path
        id := r.URL.Path[len("/events/"):]
        eventID, err := strconv.ParseUint(id, 10, 64)
        if err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }
        event, err := db.GetEventDetails(uint(eventID))
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        json.NewEncoder(w).Encode(event)
    }
}

// Update event
func UpdateEvent(db *store.Database) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var event model.Event
        if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }
        err := db.UpdateEvent(event)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        json.NewEncoder(w).Encode(event)
    }
}

// User Auth

// Place bet
