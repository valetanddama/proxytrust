# Golang: get client ip address behind proxy or balancer

[![Build Status](https://api.travis-ci.org/valetanddama/proxytrust.svg?branch=master)](https://travis-ci.org/valetanddama/proxytrust)
[![GoDoc](https://godoc.org/github.com/valetanddama/proxytrust?status.svg)](https://godoc.org/github.com/valetanddama/proxytrust)
[![Go Report Card](https://goreportcard.com/badge/github.com/valetanddama/proxytrust)](https://goreportcard.com/report/github.com/valetanddama/proxytrust)
[![codecov](https://codecov.io/gh/valetanddama/proxytrust/branch/master/graph/badge.svg)](https://codecov.io/gh/valetanddama/proxytrust)

## Installation

```
go get github.com/valetanddama/proxytrust
```

## Description
Package is suitable for those who needs to detect real client ip address if the code is on the server behind non-anonymous proxy or balancer

For detect client ip address we use X-Forwarded-For header and select left-most entry or use X-Real-IP header. If the headers are empty or invalid then you will get remote address

## Usage
```go
import "github.com/valetanddama/proxytrust"

func main() {
	routing := mux.NewRouter()
    routing = proxytrust.TrustProxyClientIp(routing)
    
    http.ListenAndServe("localhost:4000", routing)
}
```

Now, you can use **req \*http.Request** object, and get client ip address at req.RemoteAddr

## Testing
```bash
go test -race -coverprofile=coverage.txt -covermode=atomic
go test -bench=.
```

## Notice
If you are going to use package for detect client ip address at anonymous proxy or simple server, you should remember than anonymous proxy not set X-Forwarded-For header or user can set X-Forwarded-For header yourself and spoil your data
