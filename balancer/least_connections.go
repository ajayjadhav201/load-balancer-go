package balancer

import (
	"net/http"
	"net/url"
	"sync"
)

type LeastConnections struct {
	Servers []Server
	Current int
	mutex   sync.Mutex
}

func NewLeastConnections() *LeastConnections {
	return &LeastConnections{}
}

func (rr *LeastConnections) GetServers() []Server {
	return rr.Servers
}

func (rr *LeastConnections) AddServer(Url string) {
	host, err := url.Parse(Url)
	if err != nil {
		//TODO add tis error to log
		return
	}
	server := NewServer(host)
	rr.mutex.Lock()
	rr.Servers = append(rr.Servers, server)
	rr.mutex.Unlock()
}
func (rr *LeastConnections) RemoveServer(Url string) {
	index := 0
	for i, e := range rr.Servers {
		if e.GetUrl() == Url {
			index = i
		}
	}
	rr.mutex.Lock()
	rr.Servers = append(rr.Servers[:index], rr.Servers[index+1:]...) //removing a server
	rr.mutex.Unlock()
}

func (rr *LeastConnections) Serve(w http.ResponseWriter, r *http.Request) {
	//
	//TODO handle the lease connection logic
}
