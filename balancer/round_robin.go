package balancer

import (
	"net/http"
)

type RoundRobin struct {
	pool    ServerPool
	current int
	// connections int
}

func NewRoundRobin(ServerPool ServerPool) *RoundRobin {
	return &RoundRobin{pool: ServerPool}

}

func (rr *RoundRobin) Serve(w http.ResponseWriter, r *http.Request) {
	//
	poolSize := rr.pool.GetServerPoolSize()
	//
	rr.current = rr.current + 1%poolSize
	server := rr.pool.GetNextServer(rr.current)
	server.Serve(w, r)
}
