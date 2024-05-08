package balancer

import (
	"net/http"
	"net/url"
	"sync"
)

type RoundRobin struct {
	Servers []Server
	Current int
	mutex   sync.Mutex
}

func NewRoundRobin() *RoundRobin {
	return &RoundRobin{}
}

func (rr *RoundRobin) GetServers() []Server {
	return rr.Servers
}

func (rr *RoundRobin) AddServer(Url string) {
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
func (rr *RoundRobin) RemoveServer(Url string) {
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

func (rr *RoundRobin) Serve(w http.ResponseWriter, r *http.Request) {
	//
	if len(rr.Servers) == 0 {
		ServerNotAvailableWriter(w)
		return
	}
	if len(rr.Servers) == 1 {
		rr.Servers[0].GetServer().ServeHTTP(w, r) //request is sent to rpoxy server
		return
	}

	rr.Current = rr.Current + 1%len(rr.Servers)
	rr.Servers[rr.Current].GetServer().ServeHTTP(w, r) //request is sent to proxy server
}
