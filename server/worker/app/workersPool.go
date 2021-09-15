package app

import (
	"context"
	"fmt"

	"github.com/Aleksao998/LoadBalancer/worker/api/pb/loadBalancerpb"
)

func (this *App) registerToWorkersPool(port string) error {

	c := loadBalancerpb.NewLoadBalancerServiceClient(this.GrpcClient)
	req := &loadBalancerpb.RegisterRequest{
		Url: "localhost:" + port,
	}

	_, err := c.RegisterWorker(context.Background(), req)
	if err != nil {
		return fmt.Errorf("Error registering worker")
	}

	return nil
}
