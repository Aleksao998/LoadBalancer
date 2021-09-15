package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/Aleksao998/LoadBalancer/worker/api/pb/loadBalancerpb"
	"github.com/Aleksao998/LoadBalancer/worker/app"
	"google.golang.org/grpc"
)

func getPort() (*string, error) {
	portPtr := flag.String("port", "", "Port for grpc")
	flag.Parse()
	if *portPtr == "" {
		fmt.Printf("Error: Missing port from command line\n")
		return nil, fmt.Errorf("Port is missing")
	}
	return portPtr, nil
}

func main() {
	cancelChan := make(chan os.Signal, 1)
	signal.Notify(cancelChan, syscall.SIGTERM, syscall.SIGINT)

	fmt.Printf("Running Worker service\n")

	portPtr, err := getPort()
	if err != nil {
		fmt.Printf("Port is missing \n")
		return
	}

	go func(portPtr *string) {
		app := app.NewApp()
		app.Run(*portPtr)
		fmt.Printf("Exiting Load Balancer service\n")
	}(portPtr)

	sig := <-cancelChan
	fmt.Printf("Caught SIGTERM %v", sig)
	DeRegisterToWorkersPool(*portPtr)
}

func DeRegisterToWorkersPool(port string) {
	cc, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		fmt.Print(err)
	}
	c := loadBalancerpb.NewLoadBalancerServiceClient(cc)
	req := &loadBalancerpb.DeregisterRequest{
		Url: "localhost:" + port,
	}

	_, err = c.DeregisterWorker(context.Background(), req)
	if err != nil {
		fmt.Printf("Error detaching service")
	}
}
