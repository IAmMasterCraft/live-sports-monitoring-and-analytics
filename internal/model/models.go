package model

type Event struct {
    ID          uint   `json:"id" gorm:"primaryKey"`
    Name        string `json:"name"`
    Status      string `json:"status"`
    Score       string `json:"score"`
    Description string `json:"description"`
}

type User struct {
    ID       uint   `json:"id" gorm:"primaryKey"`
    Username string `json:"username"`
    Password string `json:"password"` // This should be hashed in production
}

type Bet struct {
    ID     uint   `json:"id" gorm:"primaryKey"`
    UserID uint   `json:"user_id"`
    EventID uint  `json:"event_id"`
    Stake  float64 `json:"stake"`
    Odds   float64 `json:"odds"`
}
