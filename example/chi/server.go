package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	whitelist "github.com/ihsanardanto-djoin/go-whitelist/wrapper/http"
)

func main() {
	r := chi.NewRouter()

	// Modify this for other allowed ips
	allowedIPs := []string{"127.0.0.1", "192.168.1.1"}

	// Use the middleware with the list of allowed IP addresses
	r.Use(whitelist.IPWhitelistMiddleware(allowedIPs))

	// Define your routes and handlers here
	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello, world!\n"))
	})

	r.Get("/fail", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Hello, world!\n"))
	})

	http.ListenAndServe(":8080", r)
}
