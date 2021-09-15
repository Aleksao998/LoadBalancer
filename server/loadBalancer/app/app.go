package app

import (
	"github.com/Aleksao998/LoadBalancer/api"
	"github.com/Aleksao998/LoadBalancer/services/loadBalancer"
)

func NewApp() App {
	return App{}
}

type App struct {
}

func (this *App) Run() {
	dbConnection := this.getConnection()
	loadBalancer := loadBalancer.NewLoadbalancer()

	api := &api.Api{
		Database:     dbConnection,
		LoadBalancer: loadBalancer,
	}

	go this.starHttpServer(api)
	this.startGrpc(api)
}
