package routes

import (
	"net/http"

	"github.com/pratham704/golang-ddd/internal/core/health"
	"github.com/pratham704/golang-ddd/internal/core/middleware"
	"github.com/pratham704/golang-ddd/internal/core/monitor"
	userHttp "github.com/pratham704/golang-ddd/internal/user/infrastructure/transport/http"
	"github.com/pratham704/golang-ddd/pkg/utils"
)

// SetupRoutes initializes and returns the HTTP router with all routes.
func SetupRoutes() *http.ServeMux {
	router := http.NewServeMux()

	// Register health check endpoint
	router.Handle("GET /health", middleware.LoggingMiddleware(http.HandlerFunc(health.HealthCheckHandler())))

	// Register monitoring endpoint
	router.Handle("GET /metrics", monitor.MetricsHandler())

	// register user routes
	userHttp.SetupUserRoutes(router)

	// Register a welcome message
	router.Handle("/", middleware.LimiterMiddleware(middleware.LoggingMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		utils.WriteJSONResponse(w, http.StatusOK, map[string]interface{}{"message": "Welcome to the API"})
	}))))

	return router
}
