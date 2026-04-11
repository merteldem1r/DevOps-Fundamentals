package middlewares

import (
	"fmt"
	"net"
	"net/http"
	"time"
)

func RequestLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		ip := getClientIP(r)

		fmt.Printf(
			"[REQUEST] %s | %s | %s | IP: %s | UA: %s\n",
			r.Method,
			r.URL.Path,
			start.Format(time.RFC3339),
			ip,
			r.UserAgent(),
		)

		// next handler
		next.ServeHTTP(w, r)

		fmt.Printf(
			"[DONE] %s | %s | Duration: %v\n",
			r.Method,
			r.URL.Path,
			time.Since(start),
		)

	})
}

func getClientIP(r *http.Request) string {
	// Check X-Forwarded-For first (proxy case)
	ip := r.Header.Get("X-Forwarded-For")
	if ip != "" {
		return ip
	}

	ip = r.Header.Get("X-Real-IP")
	if ip != "" {
		return ip
	}

	// fallback
	host, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return r.RemoteAddr
	}

	return host
}
