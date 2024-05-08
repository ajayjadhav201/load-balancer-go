package balancer

import (
	"net/http/httputil"
	"net/url"
)

type Server interface {
	GetServer() *httputil.ReverseProxy
	GetUrl() string
	SetAlive(bool)
	IsAlive() bool
	GetConnections() int
}

type ProxyServer struct {
	Proxy       *httputil.ReverseProxy
	Url         string
	Alive       bool
	Connections int
}

func NewServer(Url *url.URL) Server {
	proxy := httputil.NewSingleHostReverseProxy(Url)
	return &ProxyServer{
		Proxy:       proxy,
		Url:         Url.String(),
		Alive:       true,
		Connections: 0,
	}
}

func (s *ProxyServer) IsAlive() bool {
	return s.Alive
}

func (s *ProxyServer) SetAlive(status bool) {
	s.Alive = status
}

func (s *ProxyServer) GetServer() *httputil.ReverseProxy {
	return s.Proxy
}

func (s *ProxyServer) GetUrl() string {
	return s.Url
}

func (s *ProxyServer) GetConnections() int {
	return s.Connections
}
