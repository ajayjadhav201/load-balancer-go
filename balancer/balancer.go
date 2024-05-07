package balancer

import "net/http"

type LoadBalancer interface {
	Serve(http.ResponseWriter, *http.Request)
}
