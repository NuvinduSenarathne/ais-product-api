package middlewares

import (
	"ais-product-api/logger"
	"net/http"
	"time"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log := logger.Logger
		start := time.Now()
		next.ServeHTTP(w, r)
		duration := time.Since(start)

		log.Printf("[%s] %s - %s", r.Method, r.URL.Path, duration)
	})
}
