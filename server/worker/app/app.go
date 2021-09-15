package app

import (
	"fmt"
	"net"
	"time"

	"github.com/Aleksao998/LoadBalancer/worker/api"
	"github.com/Aleksao998/LoadBalancer/worker/api/pb/bankAccountpb"
	"github.com/Aleksao998/LoadBalancer/worker/api/pb/expensespb"
	"google.golang.org/grpc"
)

func NewApp() App {
	cc, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		fmt.Print(err)
	}
	return App{
		GrpcClient: cc,
	}
}

type App struct {
	GrpcClient *grpc.ClientConn
}

func (this *App) Run(port string) {
	err := this.registerToWorkersPool(port)
	if err != nil {
		time.Sleep(60 * time.Second)
		fmt.Println("started again")
		err := this.registerToWorkersPool(port)
		if err != nil {
			panic(err)
		}
	}

	lis, err := net.Listen("tcp", "0.0.0.0:"+port)
	if err != nil {
		panic("Failed to listen server")
	}
	fmt.Printf("Service starter on port: %s \n", port)

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
