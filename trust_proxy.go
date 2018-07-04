package trust_proxy

import (
	"net/http"
	"regexp"
	"strings"
)

func TrustProxyClientIp(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		regexIpPattern := "^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5]).){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])(:[0-9]+)?$"

		remoteAddr := req.RemoteAddr
		xForwardedFor := strings.TrimSpace(strings.Split(req.Header.Get("X-Forwarded-For"), ",")[0])

		if xForwardedFor != "" && regexp.MustCompile(regexIpPattern).FindString(xForwardedFor) != "" {
			remoteAddr = xForwardedFor
		}

		req.RemoteAddr = strings.Split(remoteAddr, ":")[0]
		next.ServeHTTP(res, req)
	})
}