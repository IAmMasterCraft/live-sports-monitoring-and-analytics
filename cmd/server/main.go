package main

import (
    "log"
    "fmt"
    "net/http"
	"live-event-dashboard/internal/config"
    "live-event-dashboard/internal/api"
    "live-event-dashboard/internal/store"
)

func main() {
	cfg, err := config.LoadConfig("configs")
    if err != nil {
        log.Fatalf("Failed to load configuration: %v", err)
    }
	log.Printf("Loaded configuration: %+v", cfg)

	db, err := store.NewDB(cfg.Database)
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }
	log.Printf("Connected to database: %s", cfg.Database.DBName)
	defer db.Close()

    http.HandleFunc("/", api.HandleRequests)
    
	log.Printf("Starting server on :%d", cfg.Server.Port)

    err = http.ListenAndServe(fmt.Sprintf(":%d", cfg.Server.Port), nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
