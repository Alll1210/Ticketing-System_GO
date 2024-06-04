package controllers

import (
    "encoding/json"
    "net/http"
    "ticketing-system/models"
    "ticketing-system/utils"
    "github.com/gorilla/mux"
    "strconv"
)

func GetTickets(w http.ResponseWriter, r *http.Request) {
    var tickets []models.Ticket
    utils.DB.Find(&tickets)
    json.NewEncoder(w).Encode(tickets)
}

func GetTicket(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, _ := strconv.Atoi(params["id"])
    var ticket models.Ticket
    utils.DB.First(&ticket, id)
    json.NewEncoder(w).Encode(ticket)
}

func CreateTicket(w http.ResponseWriter, r *http.Request) {
    var ticket models.Ticket
    json.NewDecoder(r.Body).Decode(&ticket)
    utils.DB.Create(&ticket)
    json.NewEncoder(w).Encode(ticket)

    // Notify all users about the new ticket
    var users []models.User
    utils.DB.Find(&users)
    for _, user := range users {
        message := "New ticket created: " + ticket.Title
        CreateNotification(user.ID, message)
    }
}

func UpdateTicket(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, _ := strconv.Atoi(params["id"])
    var ticket models.Ticket
    utils.DB.First(&ticket, id)
    json.NewDecoder(r.Body).Decode(&ticket)
    utils.DB.Save(&ticket)
    json.NewEncoder(w).Encode(ticket)

    // Notify all users about the ticket update
    var users []models.User
    utils.DB.Find(&users)
    for _, user := range users {
        message := "Ticket updated: " + ticket.Title
        CreateNotification(user.ID, message)
    }
}

func DeleteTicket(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, _ := strconv.Atoi(params["id"])
    var ticket models.Ticket
    utils.DB.Delete(&ticket, id)
    json.NewEncoder(w).Encode("The ticket is deleted successfully!")

    // Notify all users about the ticket deletion
    var users []models.User
    utils.DB.Find(&users)
    for _, user := range users {
        message := "Ticket deleted: " + ticket.Title
        CreateNotification(user.ID, message)
    }
}
