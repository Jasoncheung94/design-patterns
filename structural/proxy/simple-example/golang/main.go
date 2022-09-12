package main

import "fmt"

type Server interface {
	HandleRequest()
}

type ServerImpl struct {
}

func (s *ServerImpl) HandleRequest() {
	fmt.Println("Handling request")
}

type ProxyServer struct {
	server    Server
	rateLimit int64
}

// Proxy Server which handles a global rate limit check before the server handles the request. aka Middleware(Kinda)
func (p *ProxyServer) HandleRequest() {
	fmt.Println("Proxying request")
	if p.rateLimit > 2 {
		panic("Rate limit exceeded")
	}
	p.rateLimit++
	p.server.HandleRequest()
}

// Before the server handles the request, it's passed through a proxy server which handles a rate limit check first.
//
// GO111MODULE="off" go run main.go
func main() {
	server := &ServerImpl{}
	proxy := &ProxyServer{
		server: server,
	}
	for i := 0; i < 4; i++ {
		proxy.HandleRequest()
	}
}
