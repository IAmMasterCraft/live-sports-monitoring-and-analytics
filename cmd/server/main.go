package main

import (
    "log"
    "fmt"
    "net/http"
	"live-event-dashboard/internal/config"
    "live-event-dashboard/internal/store"
	"live-event-dashboard/pkg/middleware"
	"live-event-dashboard/internal/api"
	"live-event-dashboard/internal/service"
)

func main() {
	cfg, err := config.LoadConfig("configs")
    if err != nil {
        log.Fatalf("Failed to load configuration: %v", err)
    }

	db, err := store.NewDatabase(cfg.Database)
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }
	log.Printf("Connected to database: %s", cfg.Database.DBName)

	if err := db.AutoMigrate(); err != nil {
        log.Fatalf("Failed to migrate database schemas: %v", err)
    }

	service.GetLiveData()
    mux := api.RegisterRoutes(db)
    wrappedMux := middleware.LogRequest(middleware.ErrorHandler(mux))
    
	log.Printf("Starting server on :%d", cfg.Server.Port)

    err = http.ListenAndServe(fmt.Sprintf(":%d", cfg.Server.Port), wrappedMux)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
