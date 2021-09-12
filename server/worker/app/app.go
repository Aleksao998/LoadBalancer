package app

import (
	"net"

	"github.com/Aleksao998/LoadBalancer/worker/api"
	"github.com/Aleksao998/LoadBalancer/worker/api/pb/bankAccountpb"
	"github.com/Aleksao998/LoadBalancer/worker/api/pb/expensespb"
	"google.golang.org/grpc"
)

func NewApp() App {
	return App{}
}

type App struct {
}

type Server struct{}

func (this *App) Run() {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		panic("Failed to listen server")
	}

	s := grpc.NewServer()
	this.registerServices(s)
	if err := s.Serve(lis); err != nil {
		panic("Failed to serve server")
	}
}

func (this App) registerServices(s *grpc.Server) {
	dbConnection := this.getConnection()
	api := api.Api{
		Database: dbConnection,
	}

	bankAccountpb.RegisterBankAccountServiceServer(s, &api)
	expensespb.RegisterExpensesServiceServer(s, &api)
}
