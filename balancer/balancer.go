package balancer

import (
	"encoding/json"
	"net/http"
)

// import "net/http"

const (
	LeastConnectionSLoadBalancer BalancerType = "LC"
	RoundRobinLoadBalancer       BalancerType = "RR"
	WeightedRoundRobin           BalancerType = "WRR"
)

type BalancerType string

type Balancer interface {
	GetServers() []Server
	AddServer(string)
	RemoveServer(string)
	Serve(http.ResponseWriter, *http.Request)
}

func NewLoadBalancer(BalancerType BalancerType) Balancer {
	switch BalancerType {
	case "RR":
		return NewRoundRobin()
	}
	return nil
}

func ServerNotAvailableWriter(w http.ResponseWriter) {
	msg, _ := json.Marshal(map[string]string{"Message": "Server not available."})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusServiceUnavailable)
	w.Write(msg)
}
