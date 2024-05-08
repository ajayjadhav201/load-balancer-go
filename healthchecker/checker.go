package healthchecker

import "github.com/ajayjadhav201/load-balancer-go/balancer"

func CheckHealth(lb balancer.Balancer) {
	//

	//if server is not alive
	lb.RemoveServer("server_url_from_lb.")
}
