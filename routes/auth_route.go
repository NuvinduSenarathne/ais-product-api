package routes

import (
	"ais-product-api/controllers"

	"github.com/gorilla/mux"
)

func SetupRoutes(router *mux.Router) {
	router.HandleFunc("/register", controllers.RegisterHandler).Methods("POST")
	router.HandleFunc("/login", controllers.LoginHandler).Methods("POST")
}
