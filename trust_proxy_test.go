package trust_proxy

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestTrustProxyClientIpWithRemoteAddress(t *testing.T) {
	req, _ := http.NewRequest("GET", "/", nil)
	req.RemoteAddr = "192.168.1.1"

	testHandler := http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		if req.RemoteAddr != "192.168.1.1" {
			t.Fatal("Error")
		}
	})
	rr := httptest.NewRecorder()

	handler := TrustProxyClientIp(testHandler)
	handler.ServeHTTP(rr, req)
}

func TestTrustProxyClientIpWithRemoteAddressAndPort(t *testing.T) {
	req, _ := http.NewRequest("GET", "/", nil)
	req.RemoteAddr = "192.168.1.1:9001"

	testHandler := http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		if req.RemoteAddr != "192.168.1.1" {
			t.Fatal("Error")
		}
	})
	rr := httptest.NewRecorder()

	handler := TrustProxyClientIp(testHandler)
	handler.ServeHTTP(rr, req)
}

func TestTrustProxyClientIpWithXForwardedFor(t *testing.T) {
	req, _ := http.NewRequest("GET", "/", nil)
	req.Header.Set("X-Forwarded-For", "192.168.1.2")

	testHandler := http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		if req.RemoteAddr != "192.168.1.2" {
			t.Fatal("Error")
		}
	})
	rr := httptest.NewRecorder()

	handler := TrustProxyClientIp(testHandler)
	handler.ServeHTTP(rr, req)
}

func TestTrustProxyClientIpWithXForwardedForAndPort(t *testing.T) {
	req, _ := http.NewRequest("GET", "/", nil)
	req.Header.Set("X-Forwarded-For", "192.168.1.2:9009")

	testHandler := http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		if req.RemoteAddr != "192.168.1.2" {
			t.Fatal("Error")
		}
	})
	rr := httptest.NewRecorder()

	handler := TrustProxyClientIp(testHandler)
	handler.ServeHTTP(rr, req)
}

func TestTrustProxyClientIpWithIncorrectXForwardedFor(t *testing.T) {
	req, _ := http.NewRequest("GET", "/", nil)
	req.Header.Set("X-Forwarded-For", "incorrect_ip")
	req.RemoteAddr = "192.168.1.1"

	testHandler := http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		if req.RemoteAddr != "192.168.1.1" {
			t.Fatal("Error")
		}
	})
	rr := httptest.NewRecorder()

	handler := TrustProxyClientIp(testHandler)
	handler.ServeHTTP(rr, req)
}

func TestTrustProxyClientIpWithXRealIP(t *testing.T) {
	req, _ := http.NewRequest("GET", "/", nil)
	req.Header.Set("X-Real-IP", "192.168.1.3")

	testHandler := http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		if req.RemoteAddr != "192.168.1.3" {
			t.Fatal("Error")
		}
	})
	rr := httptest.NewRecorder()

	handler := TrustProxyClientIp(testHandler)
	handler.ServeHTTP(rr, req)
}

func TestTrustProxyClientIpWithXRealIPAndPort(t *testing.T) {
	req, _ := http.NewRequest("GET", "/", nil)
	req.Header.Set("X-Real-IP", "192.168.1.3:9002")

	testHandler := http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		if req.RemoteAddr != "192.168.1.3" {
			t.Fatal("Error")
		}
	})
	rr := httptest.NewRecorder()

	handler := TrustProxyClientIp(testHandler)
	handler.ServeHTTP(rr, req)
}

func TestTrustProxyClientIpWithIncorrectXRealIP(t *testing.T) {
	req, _ := http.NewRequest("GET", "/", nil)
	req.Header.Set("X-Real-IP", "incorrect_ip")
	req.RemoteAddr = "192.168.1.1"

	testHandler := http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		if req.RemoteAddr != "192.168.1.1" {
			t.Fatal("Error")
		}
	})
	rr := httptest.NewRecorder()

	handler := TrustProxyClientIp(testHandler)
	handler.ServeHTTP(rr, req)
}

func TestTrustProxyClientIpWithEmptyIp(t *testing.T) {
	req, _ := http.NewRequest("GET", "/", nil)
	req.RemoteAddr = ""

	testHandler := http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		if req.RemoteAddr != "" {
			t.Fatal("Error")
		}
	})
	rr := httptest.NewRecorder()

	handler := TrustProxyClientIp(testHandler)
	handler.ServeHTTP(rr, req)
}

func BenchmarkTrustProxyClientIpWithXForwardedFor(b *testing.B) {
	req, _ := http.NewRequest("GET", "/", nil)
	req.Header.Set("X-Forwarded-For", "192.168.1.2")

	testHandler := http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {})
	rr := httptest.NewRecorder()

	for i := 0; i < b.N; i++ {
		handler := TrustProxyClientIp(testHandler)
		handler.ServeHTTP(rr, req)
	}
}

func BenchmarkTrustProxyClientIpWithXRealIP(b *testing.B) {
	req, _ := http.NewRequest("GET", "/", nil)
	req.Header.Set("X-Real-IP", "192.168.1.3")

	testHandler := http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {})
	rr := httptest.NewRecorder()

	for i := 0; i < b.N; i++ {
		handler := TrustProxyClientIp(testHandler)
		handler.ServeHTTP(rr, req)
	}
}

func BenchmarkTrustProxyClientIpWithRemoteAddress(b *testing.B) {
	req, _ := http.NewRequest("GET", "/", nil)
	req.RemoteAddr = "192.168.1.2"

	testHandler := http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {})
	rr := httptest.NewRecorder()

	for i := 0; i < b.N; i++ {
		handler := TrustProxyClientIp(testHandler)
		handler.ServeHTTP(rr, req)
	}
}
