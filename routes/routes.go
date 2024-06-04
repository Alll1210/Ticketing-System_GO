package routes

import (
    "ticketing-system/controllers"
    "ticketing-system/utils"
    "github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
    router := mux.NewRouter()
    SetupUserRoutes(router)
    SetupEventRoutes(router)
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
    userRouter.HandleFunc("/bookings", controllers.GetUserBookings).Methods("GET")

    adminRouter := router.PathPrefix("/admin").Subrouter()
    adminRouter.Use(utils.AdminMiddleware)
    adminRouter.HandleFunc("/users", controllers.GetUsers).Methods("GET")
}

func SetupEventRoutes(router *mux.Router) {
    router.HandleFunc("/events", controllers.GetEvents).Methods("GET")
    router.HandleFunc("/events/{id}", controllers.GetEvent).Methods("GET")

    adminRouter := router.PathPrefix("/admin").Subrouter()
    adminRouter.Use(utils.AdminMiddleware)
    adminRouter.HandleFunc("/events", controllers.CreateEvent).Methods("POST")
    adminRouter.HandleFunc("/events/{id}", controllers.UpdateEvent).Methods("PUT")
    adminRouter.HandleFunc("/events/{id}", controllers.DeleteEvent).Methods("DELETE")

    eventRouter := router.PathPrefix("/events").Subrouter()
    eventRouter.Use(utils.AuthMiddleware)
    eventRouter.HandleFunc("/book", controllers.BookTickets).Methods("POST")
}

func SetupNotificationRoutes(router *mux.Router) {
    notificationRouter := router.PathPrefix("/notifications").Subrouter()
    notificationRouter.Use(utils.AuthMiddleware)
    notificationRouter.HandleFunc("", controllers.GetNotifications).Methods("GET")
    notificationRouter.HandleFunc("/{id}/read", controllers.MarkNotificationAsRead).Methods("PUT")
}
