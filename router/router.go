package router

import (
    "net/http"
    "yourapp/controller"
)

func SetupRoutes(controller *controller.UserController) {
    http.HandleFunc("/users", controller.Authenticate(controller.GetUsers))
    http.HandleFunc("/users/create", controller.Authenticate(controller.CreateUser))
    http.HandleFunc("/users/update", controller.Authenticate(controller.UpdateUser))
    http.HandleFunc("/users/delete", controller.Authenticate(controller.DeleteUser))
}
