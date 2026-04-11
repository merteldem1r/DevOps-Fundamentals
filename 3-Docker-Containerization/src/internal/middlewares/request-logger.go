package middlewares

import (
	"log/slog"
	"net"
	"net/http"
	"time"
)

func RequestLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		ip := getClientIP(r)

		slog.Info("request started",
			"method", r.Method,
			"path", r.URL.Path,
			"ip", ip,
			"user_agent", r.UserAgent(),
		)

		next.ServeHTTP(w, r)

		slog.Info("request finished",
			"method", r.Method,
			"path", r.URL.Path,
			"duration", time.Since(start),
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
