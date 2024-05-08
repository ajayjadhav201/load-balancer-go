package main

import (
	"log"
	"net/http"

	"github.com/ajayjadhav201/load-balancer-go/balancer"
)

func main() {
	lb := balancer.NewLoadBalancer(balancer.RoundRobinLoadBalancer)
	//
	http.HandleFunc("/api/v1", lb.Serve) //loadbalancer will route all the api request to available servers
	//
	log.Fatal(http.ListenAndServe(":80", nil))
}
