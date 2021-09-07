package main

import (
	"fmt"

	"github.com/Aleksao998/LoadBalancer/app"
)

func main() {

	fmt.Printf("Running Load Balancer service\n")

	app := app.NewApp()
	app.Run()

	fmt.Printf("Exiting Load Balancer service\n")

}
