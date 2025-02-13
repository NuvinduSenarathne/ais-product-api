package middlewares

import (
	"net/http"
	"strings"

	"ais-product-api/utils"
)

// AuthMiddleware is a middleware for JWT authentication
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Access Denied! Missing token.", http.StatusUnauthorized)
			return
		}

		// Ensure the token is split correctly
		tokenString := strings.TrimSpace(strings.Split(authHeader, "Bearer ")[1])
		if tokenString == "" {
			http.Error(w, "Access Denied! Invalid token format.", http.StatusUnauthorized)
			return
		}

		_, err := utils.ValidateToken(tokenString)
		if err != nil {
			http.Error(w, "Access Denied! Invalid token.", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
