package models

import "time"

type Notification struct {
    ID        uint      `json:"id" gorm:"primary_key"`
    UserID    uint      `json:"user_id"`
    Message   string    `json:"message"`
    Read      bool      `json:"read" gorm:"default:false"`
    CreatedAt time.Time `json:"created_at"`
}
