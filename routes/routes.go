package routes

import (
	"ais-product-api/controllers"

	"github.com/gorilla/mux"
)

// SetupRoutes defines all public routes
func SetupPublicRoutes(router *mux.Router) {
	router.HandleFunc("/api/register", controllers.RegisterHandler).Methods("POST")
	router.HandleFunc("/api/login", controllers.LoginHandler).Methods("POST")
}

// SetupProtectedRoutes defines all authenticated routes
func SetupProtectedRoutes(router *mux.Router) {
	router.HandleFunc("/api/products", controllers.GetProducts).Methods("GET")
}
