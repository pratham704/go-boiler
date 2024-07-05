package middleware

import (
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/pratham704/golang-ddd/internal/core/limiter"
	"github.com/pratham704/golang-ddd/pkg/utils"
	"golang.org/x/time/rate"
)

func LimiterMiddleware(next http.Handler) http.Handler {
	is_limit_enabled, err := strconv.ParseBool(os.Getenv("RATE_LIMIT_ENABLED"))
	if err != nil {
		is_limit_enabled = false
	}
	if !is_limit_enabled {
		return next
	}
	limit, err := strconv.Atoi(os.Getenv("RATE_LIMIT"))
	if err != nil {
		limit = 1
	}
	duration, err := time.ParseDuration(os.Getenv("RATE_LIMIT_DURATION"))
	if err != nil {
		duration = time.Second * 2
	}
	var limiter = limiter.NewIPRateLimiter(rate.Every(duration), limit)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		limiter := limiter.GetLimiter(r.RemoteAddr)
		if !limiter.Allow() {
			utils.WriteJSONError(w, http.StatusTooManyRequests, "Too many requests")
			return
		}
		next.ServeHTTP(w, r)
	})
}
