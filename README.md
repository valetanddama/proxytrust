# Go: get correct ip of client at proxy server

For usage:

```go
func main() {
	var routing = mux.NewRouter()

	http.ListenAndServe("localhost:4000", behind_proxy.DetectClientIp(routing))
}
```

And get correct ip:
```go
req.RemoteAddr
```
