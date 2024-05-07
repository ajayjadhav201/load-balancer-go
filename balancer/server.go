package balancer

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

type Server interface {
	SetAlive(bool)
	Alive() bool
	GetUrl() *url.URL
	GetActiveConnections() int
	Serve(http.ResponseWriter, http.Request)
}

func NewServer(Url *url.URL) *server {
	return &server{url: Url}
}

type server struct {
	url         *url.URL
	alive       bool
	connections int
}

func (s *server) SetAlive(Alive bool) {
	s.alive = Alive
}

func (s *server) Alive() bool {
	return s.alive
}

func (s *server) GetUrl() *url.URL {
	return s.url
}

func (s *server) GetActiveConnections() int {
	return s.connections
}

func (s *server) Serve(w http.ResponseWriter, r *http.Request) {
	Proxy := httputil.NewSingleHostReverseProxy(s.url)
	Proxy.ServeHTTP(w, r)
}
