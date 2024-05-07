package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/ajayjadhav201/golang-load-balancer/balancer"
)

var (
	authServer     = "http://localhost:8000"
	productServer1 = "http:localhost:8090"
	productServer2 = "http:localhost:8091"
)

func main() {
	//Handle Auth Server
	authUrl, err := url.Parse(authServer)
	if err != nil {
		log.Fatal("Auth url is incorrect")
	}
	AuthProxy := httputil.NewSingleHostReverseProxy(authUrl)
	http.HandleFunc("/auth", AuthProxy.ServeHTTP)
	//
	//
	sp := balancer.NewServerPool(
		balancer.NewServer(productServer1),
		balancer.NewServer(productServer2),
	)
	rr := balancer.NewRoundRobin(sp)
	// Handle Product Server
	// ProductLB := balancer.NewRRbalancer(
	// 	productServer1, productServer2)
	//
	//
	http.HandleFunc("/product", func(w http.ResponseWriter, r *http.Request) {
		rr.Serve(w, r)
		// url := ProductLB.NextServer()
		// ProductProxy := httputil.NewSingleHostReverseProxy(url)
		// ProductProxy.ServeHTTP(w, r)
	})
	// starting Load Balancer
	fmt.Print("Starting Load Balancer...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
