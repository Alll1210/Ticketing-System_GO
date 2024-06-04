package models

type Ticket struct {
    ID          uint   `json:"id" gorm:"primary_key"`
    Title       string `json:"title"`
    Description string `json:"description"`
    Status      string `json:"status"`
}
