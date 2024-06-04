package models

type Booking struct {
    ID      uint `json:"id" gorm:"primary_key"`
    UserID  uint `json:"user_id"`
    EventID uint `json:"event_id"`
    Tickets uint `json:"tickets"`
}
