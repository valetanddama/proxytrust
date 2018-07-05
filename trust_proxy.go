package trust_proxy

import (
	"net/http"
	"strings"
	"net"
)

func TrustProxyClientIp(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		remoteAddr := removePort(req.RemoteAddr)

		if xForwardedFor := req.Header.Get("X-Forwarded-For"); xForwardedFor != "" {
			xForwardedFor = removePort(strings.TrimSpace(strings.Split(xForwardedFor, ",")[0]))

			if xForwardedFor != "" && net.ParseIP(xForwardedFor) != nil {
				remoteAddr = xForwardedFor
			}
		} else if xRealIp := req.Header.Get("X-Real-IP"); xRealIp != "" {
			xRealIp = removePort(xRealIp)

			if xRealIp != "" && net.ParseIP(xRealIp) != nil {
				remoteAddr = xRealIp
			}
		}

		req.RemoteAddr = remoteAddr
		next.ServeHTTP(res, req)
	})
}

func removePort(ip string) string {
	return strings.Split(ip, ":")[0]
}