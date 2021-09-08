package app

import (
	"context"
	"fmt"
	"net"

	"github.com/Aleksao998/LoadBalancer/worker/api/pb/bankAccountpb"
	"google.golang.org/grpc"
)

func NewApp() App {
	return App{}
}

type App struct {
}

type Server struct{}

func (*Server) CreateBankAccount(ctx context.Context, req *bankAccountpb.CreateBankAccountRequest) (*bankAccountpb.CreateBankAccountResponse, error) {
	name := req.GetName()
	fmt.Printf("USAOO")
	result := "Hello " + name

	res := bankAccountpb.CreateBankAccountResponse{
		Name: result,
	}

	return &res, nil
}

func (this *App) Run() {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		panic("Failed to listen server")
	}

	s := grpc.NewServer()
	bankAccountpb.RegisterCreateBankAccountServiceServer(s, &Server{})

	if err := s.Serve(lis); err != nil {
		panic("Failed to serve server")
	}
}
