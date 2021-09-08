package main

import (
	"fmt"

	"github.com/Aleksao998/LoadBalancer/worker/app"
)

func main() {

	fmt.Printf("Running Worker service\n")

	app := app.NewApp()
	app.Run()

	fmt.Printf("Exiting Load Balancer service\n")

}
