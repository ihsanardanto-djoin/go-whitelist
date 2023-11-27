package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	ginwhitelist "github.com/ihsanardanto-djoin/go-whitelist/wrapper/gin"
)

// TODO: Move to Example/Echo
func main() {
	r := gin.Default()

	// Modify this for other allowed ips
	allowedIPs := []string{"127.0.0.2", "192.168.1.1"}

	// Use the middleware with the list of allowed IP addresses
	r.Use(ginwhitelist.IPWhitelistMiddleware(allowedIPs))

	// Define your routes and handlers here
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, World!\n")
	})

	r.Run()
}
