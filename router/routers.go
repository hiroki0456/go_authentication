package router

import (
	"net/http"
	"practice/auth/controllers"
	"practice/auth/repository"
)

func SetRouter(rep *repository.Repository) {
	http.HandleFunc("/api/users", controllers.UsersHandler(rep))
	http.HandleFunc("/api/login", controllers.AuthHandler(rep))
}
