package gowhitelist

import (
	"fmt"
	"net"
	"net/http"
	"strings"
)

type IPWhitelist struct {
	AllowedIPs []string
}

func GetClientIP(r *http.Request) (string, error) {
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
func (i *IPWhitelist) IsIPAllowed(ip string) bool {
	for _, allowedIP := range i.AllowedIPs {
		if ip == allowedIP {
			return true
		}
	}
	return false
}
