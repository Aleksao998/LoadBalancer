package api

import (
	"fmt"
	"strconv"
	"context"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/codes"
	"github.com/Aleksao998/LoadBalancer/worker/services/bankAccount"
	"github.com/Aleksao998/LoadBalancer/worker/api/pb/bankAccountpb"
)

func (this *Api) CreateBankAccount(ctx context.Context, req *bankAccountpb.BankAccountRequest) (*bankAccountpb.BankAccountResponse, error) {
	name := req.GetName()
	id := req.GetUserId()
	var result string

	ok, err := bankAccount.DoesBankAccountExists(name, id, this.Database)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal, err.Error(),
		)

	}
	if ok {
		return nil, status.Errorf(
			codes.AlreadyExists, fmt.Sprintf("User already exists"),
		)
	}

	bank_account_id, err := bankAccount.RegisterBankAccount(name,id, this.Database)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,err.Error(),
		)
	}

	result = strconv.Itoa(bank_account_id)
	res := bankAccountpb.BankAccountResponse{
		BankAccountId: result,
	}

	return &res, nil
}

func (this *Api) FetchBankAccount(ctx context.Context, req *bankAccountpb.BankAccountRequest) (*bankAccountpb.FetchBankAccountResponse, error) {
	name := req.GetName()
	id := req.GetUserId()

	ok, err := bankAccount.DoesBankAccountExists(name, id, this.Database)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal, err.Error(),
		)

	}
	if !ok {
		return nil, status.Errorf(
			codes.AlreadyExists, fmt.Sprintf("User does not exists"),
		)
	}

	bankAccountAmount, err := bankAccount.FetchBankAccountAmount(name,id, this.Database)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,err.Error(),
		)
	}


	res := bankAccountpb.FetchBankAccountResponse{
		Name: name,
		Amount:  bankAccountAmount,
	}

	return &res, nil
}

func (this *Api) DeleteBankAccount(ctx context.Context, req *bankAccountpb.BankAccountRequest) (*bankAccountpb.DeleteBankAccountResponse, error) {
	name := req.GetName()
	id := req.GetUserId()

	err := bankAccount.DeleteBankAccount(name,id, this.Database)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,err.Error(),
		)
	}

	res := bankAccountpb.DeleteBankAccountResponse{
	}

	return &res, nil
}
