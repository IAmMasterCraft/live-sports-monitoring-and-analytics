package store

import (
    "fmt"
    "gorm.io/gorm"
    "gorm.io/driver/mysql"
    "live-event-dashboard/internal/config"
)

// Database represents a GORM DB connection
type Database struct {
    *gorm.DB
}

// NewDatabase initializes a new database connection
func NewDatabase(cfg config.DatabaseConfig) (*Database, error) {
    dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
        cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DBName)
    
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        return nil, err
    }

    return &Database{db}, nil
}
