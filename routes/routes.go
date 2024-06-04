package routes

import (
    "ticketing-system/controllers"
    "ticketing-system/utils"
    "github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
    router := mux.NewRouter()
    SetupUserRoutes(router)
    SetupTicketRoutes(router)
    return router
}

func SetupUserRoutes(router *mux.Router) {
    router.HandleFunc("/register", controllers.Register).Methods("POST")
    router.HandleFunc("/login", controllers.Login).Methods("POST")
}

func SetupTicketRoutes(router *mux.Router) {
    // Rute untuk mendapatkan tiket tersedia untuk semua pengguna
    router.HandleFunc("/tickets", controllers.GetTickets).Methods("GET")
    router.HandleFunc("/tickets/{id}", controllers.GetTicket).Methods("GET")

    // Rute untuk menambah, mengedit, dan menghapus tiket hanya tersedia untuk admin
    adminRouter := router.PathPrefix("/admin").Subrouter()
    adminRouter.Use(utils.AdminMiddleware)
    adminRouter.HandleFunc("/tickets", controllers.CreateTicket).Methods("POST")
    adminRouter.HandleFunc("/tickets/{id}", controllers.UpdateTicket).Methods("PUT")
    adminRouter.HandleFunc("/tickets/{id}", controllers.DeleteTicket).Methods("DELETE")
}
