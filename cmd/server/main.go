package main

import (
    "log"
    "fmt"
    "net/http"
	"live-event-dashboard/internal/config"
    "live-event-dashboard/internal/api"
)

func main() {
	cfg, err := config.LoadConfig("configs")
    if err != nil {
        log.Fatalf("Failed to load configuration: %v", err)
    }
	log.Printf("Loaded configuration: %+v", cfg)

    http.HandleFunc("/", api.HandleRequests)
    
	log.Printf("Starting server on :%d", cfg.Server.Port)

    err = http.ListenAndServe(fmt.Sprintf(":%d", cfg.Server.Port), nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
