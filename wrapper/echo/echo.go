package echo

import (
	"net/http"

	gowhitelist "github.com/ihsanardanto-djoin/go-whitelist"
	"github.com/labstack/echo/v4"
)

func IPWhitelistMiddleware(allowedIPs []string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			clientIP, err := gowhitelist.GetClientIP(c.Request())
			if err != nil {
				return c.JSON(http.StatusUnauthorized, "Unauthorized, IP Client not detected\n")
			}

			clientWList := gowhitelist.IPWhitelist{AllowedIPs: allowedIPs}

			// Check if the client's IP is in the whitelist
			if !clientWList.IsIPAllowed(clientIP) {
				return c.String(http.StatusForbidden, "Forbidden\n")
			}

			return next(c)
		}
	}
}
