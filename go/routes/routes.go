package routes

import (
	"github.com/gorilla/mux"
	"golang/start/go/controllers"
)

func InitRoutes() *mux.Router {
	r := mux.NewRouter()

	// Public routes
	r.HandleFunc("/signup", controllers.SignupHandler).Methods("POST")
	r.HandleFunc("/login", controllers.LoginHandler).Methods("POST")

	// You can add protected routes here, using middleware to protect them

	return r
}
