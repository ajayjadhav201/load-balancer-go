package balancer

import "net/http"

type LeastConnections struct {
	pool    ServerPool
	current int
}

func NewLeastConnections(ServerPool ServerPool) *LeastConnections {
	return &LeastConnections{pool: ServerPool}
}

func (lc *LeastConnections) Serve(w http.ResponseWriter, r *http.Request) {
	// TODO
	server := lc.pool.GetNextServer(lc.current)
	server.Serve(w, r)
}
