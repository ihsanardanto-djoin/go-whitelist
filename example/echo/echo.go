package main

import (
	"net/http"

	"github.com/labstack/echo/v4"

	echowhitelist "github.com/ihsanardanto-djoin/go-whitelist/wrapper/echo"
)

// TODO: Move to Example/Echo
func main() {
	e := echo.New()

	// Modify this for other allowed ips
	allowedIPs := []string{"127.0.0.1", "192.168.1.1"}

	// Use the middleware with the list of allowed IP addresses
	e.Use(echowhitelist.IPWhitelistMiddleware(allowedIPs))

	// Define your routes and handlers here
	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!\n")
	})

	e.Start(":8080")
}
