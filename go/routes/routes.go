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
	r.HandleFunc("/upload", controllers.UploadHandler).Methods("POST")

	return r
}
