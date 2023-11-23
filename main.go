package main

import (
	"fmt"
	"net"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type IPWhitelist struct {
	db string
	// ipList []string
}

func getClientIP(r *http.Request) (string, error) {
	//Get IP from Cloudflare
	ip := r.Header.Get("CF-Connecting-IP")
	netIP := net.ParseIP(ip)
	if netIP != nil {
		return ip, nil
	}

	//Get IP from the X-REAL-IP header
	ip = r.Header.Get("X-REAL-IP")
	netIP = net.ParseIP(ip)
	if netIP != nil {
		return ip, nil
	}

	//Get IP from X-FORWARDED-FOR header
	ips := r.Header.Get("X-FORWARDED-FOR")
	splitIps := strings.Split(ips, ",")
	for _, ip := range splitIps {
		netIP := net.ParseIP(ip)
		if netIP != nil {
			return ip, nil
		}
	}

	//Get IP from RemoteAddr
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return "", err
	}
	netIP = net.ParseIP(ip)
	if netIP != nil {
		return ip, nil
	}
	return "", fmt.Errorf("No valid ip found")
}

// isIPAllowed checks if an IP is in the whitelist
func (i *IPWhitelist) isIPAllowed(ip string) bool {
	allowedIPs := i.getListIP()
	for _, allowedIP := range allowedIPs {
		if ip == allowedIP {
			return true
		}
	}
	return false
}

func (i *IPWhitelist) getListIP() []string {
	// TODO: Change to get from db
	return []string{"127.0.0.1", "192.168.1.1"}
}

// TODO: Move to Wrapper/Echo
func IPWhitelistMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			clientIP, err := getClientIP(c.Request())
			fmt.Println("test: ", clientIP)
			if err != nil {
				return c.JSON(http.StatusUnauthorized, "Unauthorized, IP Client not detected\n")
			}

			clientWList := IPWhitelist{db: "db_name"}

			// Check if the client's IP is in the whitelist
			if !clientWList.isIPAllowed(clientIP) {
				return c.String(http.StatusForbidden, "Forbidden\n")
			}

			// Call the next handler in the chain
			return next(c)
		}
	}
}

// TODO: Move to Example/Echo
func main() {
	// Create a new Echo instance
	e := echo.New()

	// Use the middleware with the list of allowed IP addresses
	e.Use(IPWhitelistMiddleware())

	// Define your routes and handlers here
	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!\n")
	})
	// Start the server
	e.Start(":8080")
}
