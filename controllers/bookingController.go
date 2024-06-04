package controllers

import (
    "encoding/json"
    "net/http"
    "log"
    "ticketing-system/models"
    "ticketing-system/utils"
)

func BookTickets(w http.ResponseWriter, r *http.Request) {
    userID, ok := r.Context().Value(utils.ContextKeyUserID).(uint)
    if !ok {
        http.Error(w, "User ID not found in context", http.StatusInternalServerError)
        return
    }

    var booking models.Booking
    json.NewDecoder(r.Body).Decode(&booking)

    log.Printf("Booking request: %+v\n", booking)
    booking.UserID = userID

    var event models.Event
    if err := utils.DB.First(&event, booking.EventID).Error; err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    if event.Tickets < booking.Tickets {
        http.Error(w, "Not enough tickets available", http.StatusBadRequest)
        return
    }

    event.Tickets -= booking.Tickets
    utils.DB.Save(&event)
    if err := utils.DB.Create(&booking).Error; err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    log.Printf("Booking created: %+v\n", booking)

    json.NewEncoder(w).Encode(booking)
}

func GetUserBookings(w http.ResponseWriter, r *http.Request) {
    userID, ok := r.Context().Value(utils.ContextKeyUserID).(uint)
    if !ok {
        http.Error(w, "User ID not found in context", http.StatusInternalServerError)
        return
    }
    log.Printf("Fetching bookings for user ID: %d\n", userID)

    var bookings []models.Booking
    if err := utils.DB.Where("user_id = ?", userID).Find(&bookings).Error; err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    log.Printf("Bookings found: %+v\n", bookings)

    var events []models.Event
    for _, booking := range bookings {
        var event models.Event
        if err := utils.DB.First(&event, booking.EventID).Error; err == nil {
            events = append(events, event)
        }
    }

    json.NewEncoder(w).Encode(events)
}
