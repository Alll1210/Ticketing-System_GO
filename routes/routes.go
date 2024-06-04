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
    SetupNotificationRoutes(router)
    return router
}

func SetupUserRoutes(router *mux.Router) {
    router.HandleFunc("/register", controllers.Register).Methods("POST")
    router.HandleFunc("/login", controllers.Login).Methods("POST")

    userRouter := router.PathPrefix("/user").Subrouter()
    userRouter.Use(utils.AuthMiddleware)
    userRouter.HandleFunc("/profile", controllers.GetUserProfile).Methods("GET")
    userRouter.HandleFunc("/profile", controllers.UpdateUserProfile).Methods("PUT")
}

func SetupTicketRoutes(router *mux.Router) {
    // Rute untuk mendapatkan daftar tiket dan melihat detail tiket
    router.HandleFunc("/tickets", controllers.GetTickets).Methods("GET")
    router.HandleFunc("/tickets/{id}", controllers.GetTicket).Methods("GET")

    // Rute untuk menambah, mengedit, dan menghapus tiket hanya tersedia untuk admin
    adminRouter := router.PathPrefix("/admin").Subrouter()
    adminRouter.Use(utils.AdminMiddleware)
    adminRouter.HandleFunc("/tickets", controllers.CreateTicket).Methods("POST")
    adminRouter.HandleFunc("/tickets/{id}", controllers.UpdateTicket).Methods("PUT")
    adminRouter.HandleFunc("/tickets/{id}", controllers.DeleteTicket).Methods("DELETE")
}

func SetupNotificationRoutes(router *mux.Router) {
    notificationRouter := router.PathPrefix("/notifications").Subrouter()
    notificationRouter.Use(utils.AuthMiddleware)
    notificationRouter.HandleFunc("", controllers.GetNotifications).Methods("GET")
    notificationRouter.HandleFunc("/{id}/read", controllers.MarkNotificationAsRead).Methods("PUT")
}
