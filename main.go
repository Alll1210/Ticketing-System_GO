package main

import (
    "log"
    "net/http"
    "ticketing-system/routes"
    "github.com/joho/godotenv"
)

func main() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    router := routes.SetupRouter()
    log.Fatal(http.ListenAndServe(":8080", router))
}
