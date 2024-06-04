package routes

import (
    "ticketing-system/controllers"
    "github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
    router := mux.NewRouter()
    SetupUserRoutes(router)
    SetupTicketRoutes(router)  // Pastikan fungsi ini didefinisikan
    return router
}

func SetupUserRoutes(router *mux.Router) {
    router.HandleFunc("/register", controllers.Register).Methods("POST")
    router.HandleFunc("/login", controllers.Login).Methods("POST")
}

func SetupTicketRoutes(router *mux.Router) {
    router.HandleFunc("/tickets", controllers.GetTickets).Methods("GET")
    router.HandleFunc("/tickets/{id}", controllers.GetTicket).Methods("GET")
    router.HandleFunc("/tickets", controllers.CreateTicket).Methods("POST")
    router.HandleFunc("/tickets/{id}", controllers.UpdateTicket).Methods("PUT")
    router.HandleFunc("/tickets/{id}", controllers.DeleteTicket).Methods("DELETE")
}
