package behind_proxy

import (
	"net/http"
)

func init() {}

func detectClientIp(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		remoteAddr := req.RemoteAddr
		xForwardedFor := req.Header.Get("X-Forwarded-For")

		if xForwardedFor != "" {
			remoteAddr = xForwardedFor
		}

		req.RemoteAddr = remoteAddr
		next.ServeHTTP(res, req)
	})
}
