package store

import (
    "fmt"
    "live-event-dashboard/internal/config"
    "github.com/jmoiron/sqlx"
    _ "github.com/go-sql-driver/mysql" // MySQL driver
)

// DB stores the database session information. Needs to be initialized once
type DB struct {
    *sqlx.DB
}

// NewDB creates a new database connection to the MySQL server
func NewDB(cfg config.DatabaseConfig) (*DB, error) {
    dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true",
        cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DBName)
    
    db, err := sqlx.Connect("mysql", dsn)
    if err != nil {
        return nil, err
    }

    return &DB{db}, nil
}
