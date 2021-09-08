package api

import (
	"fmt"
	"net/http"

	"google.golang.org/grpc"
)

func (this Api) CreateBankAccount(w http.ResponseWriter, r *http.Request) {
	//specify status code
	w.WriteHeader(http.StatusOK)
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		fmt.Printf("ERror")
	}
	c := bankAccountpb.CreateBankAccount()
	//update response writer
	fmt.Fprintf(w, "Api for CreateBankAccount")
}

func (this Api) FetchBankAccount(w http.ResponseWriter, r *http.Request) {
	//specify status code
	w.WriteHeader(http.StatusOK)

	//update response writer
	fmt.Fprintf(w, "Api for FetchBankAccount")
}

func (this Api) DeleteBankAccount(w http.ResponseWriter, r *http.Request) {
	//specify status code
	w.WriteHeader(http.StatusOK)

	//update response writer
	fmt.Fprintf(w, "Api for DeleteBankAccount")
}
