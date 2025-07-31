package controller

import (
    "net/http"
    "encoding/json"
    "os"
    "strings"
    "strconv"
    "github.com/golang-jwt/jwt/v5"
    "yourapp/service"
    "yourapp/model"
)

type UserController struct {
    Service *service.UserService
}

// JWT authentication middleware
func (c *UserController) Authenticate(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        tokenString := r.Header.Get("Authorization")
        if tokenString == "" {
            http.Error(w, "Missing Authorization header", http.StatusUnauthorized)
            return
        }
        tokenString = strings.TrimPrefix(tokenString, "Bearer ")

        decryptionKey := r.Header.Get("X-JWT-KEY")
        if decryptionKey == "" {
            http.Error(w, "Missing X-JWT-KEY header", http.StatusUnauthorized)
            return
        }

        // Use the provided key to validate the JWT (HMAC example)
        token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
            // Optionally, combine with private key from env
            privateKey := os.Getenv("JWT_PRIVATE_KEY")
            return []byte(decryptionKey + privateKey), nil
        })
        if err != nil || !token.Valid {
            http.Error(w, "Invalid token", http.StatusUnauthorized)
            return
        }
        next(w, r)
    }
}

func (c *UserController) GetUsers(w http.ResponseWriter, r *http.Request) {
    users, err := c.Service.GetUsers()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(users)
}

func (c *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
    var user model.User
    json.NewDecoder(r.Body).Decode(&user)
    err := c.Service.CreateUser(user)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusCreated)
}

func (c *UserController) UpdateUser(w http.ResponseWriter, r *http.Request) {
    var user model.User
    json.NewDecoder(r.Body).Decode(&user)
    err := c.Service.UpdateUser(user)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusOK)
}

func (c *UserController) DeleteUser(w http.ResponseWriter, r *http.Request) {
    idStr := r.URL.Query().Get("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        http.Error(w, "Invalid ID", http.StatusBadRequest)
        return
    }

    err = c.Service.DeleteUser(id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusOK)
}
