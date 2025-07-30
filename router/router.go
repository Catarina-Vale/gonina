package router

import (
    "net/http"
    "yourapp/controller"
)

func SetupRoutes(controller *controller.UserController) {
    http.HandleFunc("/users", controller.GetUsers)
    http.HandleFunc("/users/create", controller.CreateUser)
    http.HandleFunc("/users/update", controller.UpdateUser)
    http.HandleFunc("/users/delete", controller.DeleteUser)
}
