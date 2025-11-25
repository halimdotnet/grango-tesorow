package middleware

import (
	"encoding/json"
	"log"
	"net"
	"net/http"
	"strconv"
	"strings"

	"github.com/halimdotnet/grango-tesorow/internal/pkg/constants"
	"github.com/ulule/limiter/v3"
	"github.com/ulule/limiter/v3/drivers/store/memory"
)

func RateLimiter(next http.Handler) http.Handler {

	rate := limiter.Rate{
		Period: constants.RatelimitPeriod,
		Limit:  constants.RateLimitAttempt,
	}

	store := memory.NewStore()
	ipLimiter := limiter.New(store, rate)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip := clientIP(r)
		if ip == "" {
			ip = r.RemoteAddr
		}

		context, err := ipLimiter.Get(r.Context(), ip)
		if err != nil {
			log.Printf(
				"Rate limiter error from IP: %v, %s on %s",
				err,
				ip,
				r.URL,
			)

			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			w.WriteHeader(http.StatusServiceUnavailable)

			if err = json.NewEncoder(w).Encode(map[string]interface{}{
				"error":   true,
				"message": "Service temporarily unavailable",
			}); err != nil {
				log.Printf(
					"Failed to encode rate limiter error response for IP %s: %v",
					ip,
					err,
				)
			}

			return
		}

		w.Header().Set(
			"Content-Type",
			"application/json; charset=utf-8",
		)
		w.Header().Set(
			"X-RateLimit-Limit",
			strconv.FormatInt(context.Limit, 10),
		)
		w.Header().Set(
			"X-RateLimit-Remaining",
			strconv.FormatInt(context.Remaining, 10),
		)
		w.Header().Set(
			"X-RateLimit-Reset",
			strconv.FormatInt(context.Reset, 10),
		)

		if context.Reached {
			log.Printf(
				"Too Many Requests from %s on %s",
				ip,
				r.URL,
			)

			w.WriteHeader(http.StatusTooManyRequests)

			if err := json.NewEncoder(w).Encode(map[string]interface{}{
				"error":   true,
				"message": "Rate limit exceeded",
			}); err != nil {
				log.Printf("Failed to encode rate limit response: %v", err)
				log.Printf(
					"Failed to encode rate limit response from %s on %s",
					ip,
					r.URL,
				)
			}

			return
		}

		next.ServeHTTP(w, r)

	})
}

func clientIP(r *http.Request) string {
	var ip string

	if tcip := r.Header.Get(http.CanonicalHeaderKey("True-Client-IP")); tcip != "" {
		ip = tcip
	} else if xrip := r.Header.Get(http.CanonicalHeaderKey("X-Real-IP")); xrip != "" {
		ip = xrip
	} else if xff := r.Header.Get(http.CanonicalHeaderKey("X-Forwarded-For")); xff != "" {
		ip, _, _ = strings.Cut(xff, ",")
	}
	if ip == "" || net.ParseIP(ip) == nil {
		return ""
	}
	return ip
}
