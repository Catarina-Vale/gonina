package controller

import (
    "net/http"
    "encoding/json"
    "yourapp/service"
)

type UserController struct {
    Service *service.UserService
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
