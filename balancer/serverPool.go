package balancer

import "log"

type ServerPool interface {
	GetServers() []*server
	AddServer(*server)
	GetServerPoolSize() int
	GetNextServer(int) *server
}

type Pool struct {
	servers []*server
	size    int
}

func NewServerPool(newServers ...*server) *Pool {
	if len(newServers) == 0 {
		log.Fatal("Add atlease one server")
	}
	p := &Pool{}
	//
	p.servers = append(p.servers, newServers...)
	return p
}

func (p *Pool) GetServers() []*server {
	return p.servers
}

func (p *Pool) AddServer(Server *server) {
	p.servers = append(p.servers, Server)
}

func (p *Pool) GetServerPoolSize() int {
	return p.size
}
func (p *Pool) GetNextServer(index int) *server {
	return p.servers[index]
}
