package controllers

import (
    "encoding/json"
    "net/http"
    "ticketing-system/models"
    "ticketing-system/utils"
    "github.com/gorilla/mux"
    "strconv"
)

func GetNotifications(w http.ResponseWriter, r *http.Request) {
    userID := r.Context().Value("userID").(uint)
    var notifications []models.Notification
    utils.DB.Where("user_id = ?", userID).Find(&notifications)
    json.NewEncoder(w).Encode(notifications)
}

func CreateNotification(userID uint, message string) {
    notification := models.Notification{
        UserID:  userID,
        Message: message,
    }
    utils.DB.Create(&notification)
}

func MarkNotificationAsRead(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, _ := strconv.Atoi(params["id"])
    userID := r.Context().Value("userID").(uint)

    var notification models.Notification
    utils.DB.Where("id = ? AND user_id = ?", id, userID).First(&notification)
    notification.Read = true
    utils.DB.Save(&notification)

    json.NewEncoder(w).Encode(notification)
}
