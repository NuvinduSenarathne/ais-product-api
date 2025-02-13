package main

import (
	"net/http"
	"os"

	"ais-product-api/config"
	"ais-product-api/logger"
	"ais-product-api/middlewares"
	"ais-product-api/routes"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	// Initialize global logger
	logger.InitLogger()
	log := logger.Logger

	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Connect to Database
	config.ConnectDB()

	// Initialize Router
	router := mux.NewRouter()

	// Apply Middlewares
	router.Use(middlewares.LoggingMiddleware)
	router.Use(middlewares.CORSMiddleware)

	// Setup Public Routes
	routes.SetupPublicRoutes(router)

	// Protected routes (require authentication)
	protectedRoutes := router.PathPrefix("/api").Subrouter()
	protectedRoutes.Use(middlewares.AuthMiddleware)
	routes.SetupProtectedRoutes(protectedRoutes)

	// Start Server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Println("Server running on port", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
