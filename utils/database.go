package utils

import (
    "fmt"
    "log"
    "os"
    "ticketing-system/models"
    "golang.org/x/crypto/bcrypt"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "github.com/joho/godotenv"
)

var DB *gorm.DB

func init() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    dsn := fmt.Sprintf(
        "host=%s user=%s password=%s dbname=%s port=5432 sslmode=require",
        os.Getenv("DB_HOST"),
        os.Getenv("DB_USER"),
        os.Getenv("DB_PASSWORD"),
        os.Getenv("DB_NAME"),
    )

    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database!")
    }

    DB.AutoMigrate(&models.User{}, &models.Event{}, &models.Booking{}, &models.Notification{})

    createAdminIfNotExists()
}

func createAdminIfNotExists() {
    var admin models.User
    result := DB.Where("email = ?", "admin@example.com").First(&admin)
    if result.Error != nil && result.Error == gorm.ErrRecordNotFound {
        hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("adminpassword"), bcrypt.DefaultCost)
        admin = models.User{
            Username: "admin",
            Email:    "admin@example.com",
            Password: string(hashedPassword),
            Role:     "admin",
        }
        DB.Create(&admin)
        log.Println("Admin user created")
    }
}
