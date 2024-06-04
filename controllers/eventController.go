package controllers

import (
    "encoding/json"
    "net/http"
    "ticketing-system/models"
    "ticketing-system/utils"
    "github.com/gorilla/mux"
    "strconv"
)

func GetEvents(w http.ResponseWriter, r *http.Request) {
    var events []models.Event
    utils.DB.Find(&events)
    json.NewEncoder(w).Encode(events)
}

func GetEvent(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, _ := strconv.Atoi(params["id"])
    var event models.Event
    utils.DB.First(&event, id)
    json.NewEncoder(w).Encode(event)
}

func CreateEvent(w http.ResponseWriter, r *http.Request) {
    var event models.Event
    json.NewDecoder(r.Body).Decode(&event)
    utils.DB.Create(&event)
    json.NewEncoder(w).Encode(event)

    // Notify all users about the new event
    var users []models.User
    utils.DB.Find(&users)
    for _, user := range users {
        message := "New event created: " + event.Title
        CreateNotification(user.ID, message)
    }
}

func UpdateEvent(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, _ := strconv.Atoi(params["id"])
    var event models.Event
    utils.DB.First(&event, id)
    json.NewDecoder(r.Body).Decode(&event)
    utils.DB.Save(&event)
    json.NewEncoder(w).Encode(event)

    // Notify all users about the event update
    var users []models.User
    utils.DB.Find(&users)
    for _, user := range users {
        message := "Event updated: " + event.Title
        CreateNotification(user.ID, message)
    }
}

func DeleteEvent(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, _ := strconv.Atoi(params["id"])
    var event models.Event
    utils.DB.Delete(&event, id)
    json.NewEncoder(w).Encode("The event is deleted successfully!")

    // Notify all users about the event deletion
    var users []models.User
    utils.DB.Find(&users)
    for _, user := range users {
        message := "Event deleted: " + event.Title
        CreateNotification(user.ID, message)
    }
}
