package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"ais-product-api/config"
	"ais-product-api/models"
	"ais-product-api/utils"

	"golang.org/x/crypto/bcrypt"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}

	_, err = config.DB.Exec("INSERT INTO user (username, password) VALUES (?, ?)", user.Username, hashedPassword)
	if err != nil {
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User registered"})
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Fetch user from the database
	var storedUser models.User
	err = config.DB.QueryRow("SELECT id, username, password FROM user WHERE username = ?", user.Username).
		Scan(&storedUser.ID, &storedUser.Username, &storedUser.Password)

	if err == sql.ErrNoRows {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	} else if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	// Debugging: Print password hashes to compare
	fmt.Println("Stored Hashed Password:", storedUser.Password)
	fmt.Println("User Entered Password:", user.Password)

	// Generate JWT Token
	token, err := utils.GenerateToken(storedUser.Username)
	if err != nil {
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}

	// Return the token
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}

func GetProducts(w http.ResponseWriter, r *http.Request) {
	// Example response (replace with actual product fetching logic)
	products := []string{"Product1", "Product2", "Product3"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}
