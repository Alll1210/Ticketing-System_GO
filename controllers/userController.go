package controllers

import (
    "encoding/json"
    "net/http"
    "ticketing-system/models"
    "ticketing-system/utils"
    "golang.org/x/crypto/bcrypt"
    "gorm.io/gorm"
)

func Register(w http.ResponseWriter, r *http.Request) {
    var user models.User
    json.NewDecoder(r.Body).Decode(&user)

    // Set role to "user" by default
    user.Role = "user"

    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    user.Password = string(hashedPassword)

    if err := utils.DB.Create(&user).Error; err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(user)
}

func Login(w http.ResponseWriter, r *http.Request) {
    var user models.User
    var dbUser models.User
    json.NewDecoder(r.Body).Decode(&user)

    if err := utils.DB.Where("email = ?", user.Email).First(&dbUser).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            http.Error(w, "Invalid email or password", http.StatusUnauthorized)
        } else {
            http.Error(w, err.Error(), http.StatusInternalServerError)
        }
        return
    }

    if err := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password)); err != nil {
        http.Error(w, "Invalid email or password", http.StatusUnauthorized)
        return
    }

    token, err := utils.GenerateJWT(dbUser)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(map[string]string{"token": token})
}
