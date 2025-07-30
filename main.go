package main

import (
    "yourapp/config"
    "yourapp/controller"
    "yourapp/service"
    "yourapp/repository"
    "yourapp/router"
)

func main() {
    db := config.ConnectDB()
    userRepo := &repository.UserRepo{DB: db}
    userService := &service.UserService{Repo: userRepo}
    userController := &controller.UserController{Service: userService}

    router.SetupRoutes(userController)
    http.ListenAndServe(":8080", nil)
}
