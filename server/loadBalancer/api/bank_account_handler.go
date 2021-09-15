package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Aleksao998/LoadBalancer/api/pb/bankAccountpb"
)

type BankAccount struct {
	Name   string `json:"bank_account_name"`
	UserId string `json:"bank_account_user_id"`
}

type BankAccountOutput struct {
	BankAccountId string `json:"bank_account_id"`
}

type FetchBankAccountOutput struct {
	Name   string  `json:"bank_account_name"`
	Amount float32 `json:"bank_account_amount"`
}

func (this *Api) CreateBankAccount(w http.ResponseWriter, r *http.Request) {
	var bankAccount BankAccount
	json.NewDecoder(r.Body).Decode(&bankAccount)

	worker := this.LoadBalancer.Next()
	if worker == nil {
		this.responseError(w, 400, fmt.Errorf("No Worker Available"))
		return
	}

	req := &bankAccountpb.BankAccountRequest{
		Name:   bankAccount.Name,
		UserId: bankAccount.UserId,
	}

	res, err := worker.BankService.CreateBankAccount(context.Background(), req)
	if err != nil {
		this.responseError(w, 400, err)
		return
	}

	responseJson(w, BankAccountOutput{BankAccountId: res.BankAccountId})
}

func (this Api) FetchBankAccount(w http.ResponseWriter, r *http.Request) {
	var bankAccount BankAccount
	json.NewDecoder(r.Body).Decode(&bankAccount)

	worker := this.LoadBalancer.Next()
	if worker == nil {
		this.responseError(w, 400, fmt.Errorf("No Worker Available"))
		return
	}

	req := &bankAccountpb.BankAccountRequest{
		Name:   bankAccount.Name,
		UserId: bankAccount.UserId,
	}

	res, err := worker.BankService.FetchBankAccount(context.Background(), req)
	if err != nil {
		this.responseError(w, 400, err)
		return
	}

	responseJson(w, FetchBankAccountOutput{Name: res.Name, Amount: res.Amount})
}

func (this Api) DeleteBankAccount(w http.ResponseWriter, r *http.Request) {
	var bankAccount BankAccount
	json.NewDecoder(r.Body).Decode(&bankAccount)

	worker := this.LoadBalancer.Next()
	if worker == nil {
		this.responseError(w, 400, fmt.Errorf("No Worker Available"))
		return
	}

	req := &bankAccountpb.BankAccountRequest{
		Name:   bankAccount.Name,
		UserId: bankAccount.UserId,
	}

	_, err := worker.BankService.DeleteBankAccount(context.Background(), req)
	if err != nil {
		this.responseError(w, 400, err)
		return
	}

	responseJson(w, ApiSuccess{Status: "success"})
}
