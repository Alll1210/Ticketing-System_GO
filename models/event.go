package models

type Event struct {
    ID          uint   `json:"id" gorm:"primary_key"`
    Title       string `json:"title"`
    Description string `json:"description"`
    Date        string `json:"date"`
    Location    string `json:"location"`
    Tickets     uint   `json:"-"`
}
