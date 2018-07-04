# Go lang: get client ip address behind proxy or balancer

## Installation

```
go get github.com/valetanddama/trust-proxy
```

## Description
Package is suitable for those who need to detect real client ip address if the code is on the server behind non-anonymous proxy or balancer

For detect client ip address we use X-Forwarded-For header and select left-most entry or use X-Real-IP header. If the headers is empty or invalid then you will get remote address

## Usage
```go
import "github.com/valetanddama/trust-proxy"

func main() {
    var routing = mux.NewRouter()
    http.ListenAndServe("localhost:4000", trust_proxy.TrustProxyClientIp(routing))
}
```

Now, you can use **req \*http.Request** object, and get client ip address at req.RemoteAddr

## Notice
If you are going to use package for detect client ip address at anonymous proxy or simple server, you should remember than anonymous proxy not set X-Forwarded-For header or user can set X-Forwarded-For header yourself and spoil your data