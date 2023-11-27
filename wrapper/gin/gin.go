package gin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	gowhitelist "github.com/ihsanardanto-djoin/go-whitelist"
)

func IPWhitelistMiddleware(allowedIPs []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		clientIP, err := gowhitelist.GetClientIP(c.Request)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, "Unauthorized, IP Client not detected\n")
			return
		}

		clientWList := gowhitelist.IPWhitelist{AllowedIPs: allowedIPs}

		// Check if the client's IP is in the whitelist
		if !clientWList.IsIPAllowed(clientIP) {
			c.AbortWithStatusJSON(http.StatusForbidden, "Forbidden\n")
			return
		}
		c.Next()

	}
}
