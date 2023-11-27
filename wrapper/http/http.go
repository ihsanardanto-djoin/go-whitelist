package http

import (
	"net/http"

	gowhitelist "github.com/ihsanardanto-djoin/go-whitelist"
)

func IPWhitelistMiddleware(allowedIPs []string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			var rw *responseWriter

			clientIP, err := gowhitelist.GetClientIP(r)
			if err != nil {
				rw = &responseWriter{w, http.StatusUnauthorized, "Unauthorized, IP Client not detected\n"}
				return
			}

			clientWList := gowhitelist.IPWhitelist{AllowedIPs: allowedIPs}

			// Check if the client's IP is in the whitelist
			if !clientWList.IsIPAllowed(clientIP) {
				rw = &responseWriter{w, http.StatusForbidden, "Forbidden\n"}
				return
			}

			// Serve the next handler
			next.ServeHTTP(rw, r)
		})
	}
}

// Custom responseWriter to capture status code
type responseWriter struct {
	http.ResponseWriter
	status  int
	message string
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.status = code
	rw.ResponseWriter.WriteHeader(code)
}
