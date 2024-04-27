package store

import (
    "fmt"
    "gorm.io/gorm"
    "gorm.io/driver/mysql"
    "live-event-dashboard/internal/config"
    "live-event-dashboard/internal/model"
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

func (db *Database) AutoMigrate() error {
    // Migrate the schema for each model
    err := db.Migrator().AutoMigrate(&model.Event{})
    if err != nil {
        return err
    }
    err = db.Migrator().AutoMigrate(&model.User{})
    if err != nil {
        return err
    }
    return db.Migrator().AutoMigrate(&model.Bet{})
}


func (db *Database) GetLiveEvents() ([]model.Event, error) {
    var events []model.Event
    result := db.Find(&events) // Assuming there is some flag or status indicating 'live'
    return events, result.Error
}

func (db *Database) GetEventDetails(id uint) (model.Event, error) {
    var event model.Event
    result := db.First(&event, id)
    return event, result.Error
}

func (db *Database) UpdateEvent(event model.Event) error {
    return db.Save(&event).Error
}