package app

import (
	"fmt"
	"net"

	"github.com/Aleksao998/LoadBalancer/api"
	"github.com/Aleksao998/LoadBalancer/api/pb/loadBalancerpb"
	"google.golang.org/grpc"
)

func (this App) startGrpc(api *api.Api) {
	lis, err := net.Listen("tcp", "0.0.0.0:50052")
	if err != nil {
		panic("Failed to listen server")
	}

	s := grpc.NewServer()
	this.registerServices(s, api)

	fmt.Println("start Grpc")
	if err := s.Serve(lis); err != nil {
		panic("Failed to serve server")
	}
}

func (this App) registerServices(s *grpc.Server, api *api.Api) {
	loadBalancerpb.RegisterLoadBalancerServiceServer(s, api)
}
