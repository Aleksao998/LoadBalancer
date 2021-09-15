package api

import (
	"context"
	"fmt"

	"github.com/Aleksao998/LoadBalancer/api/pb/loadBalancerpb"
	"github.com/Aleksao998/LoadBalancer/services/loadBalancer"
)

func (this Api) printWorkerPool(workerPool []*loadBalancer.WorkerService) {
	fmt.Println("WorkerPool:")
	for _, worker := range this.LoadBalancer.Workers {
		fmt.Println(worker.Url + " ")
	}
	fmt.Println("-----------")
}

func (this *Api) RegisterWorker(ctx context.Context, req *loadBalancerpb.RegisterRequest) (*loadBalancerpb.RegisterResponse, error) {
	urlInput := req.GetUrl()

	this.LoadBalancer.AddUrl(urlInput)
	this.printWorkerPool(this.LoadBalancer.Workers)

	res := loadBalancerpb.RegisterResponse{}
	return &res, nil
}

func (this *Api) DeregisterWorker(ctx context.Context, req *loadBalancerpb.DeregisterRequest) (*loadBalancerpb.DeregisterResponse, error) {
	urlInput := req.GetUrl()

	this.LoadBalancer.RemoveUrl(urlInput)
	this.printWorkerPool(this.LoadBalancer.Workers)

	res := loadBalancerpb.DeregisterResponse{}
	return &res, nil
}
