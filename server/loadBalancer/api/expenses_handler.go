package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Aleksao998/LoadBalancer/api/pb/expensespb"
)

type Expense struct {
	Name          string  `json:"expense_name"`
	Amount        float32 `json:"expense_amount"`
	BankAccountId string  `json:"expense_bank_account_id"`
}

type ExpenseOutput struct {
	ExpenseId string `json:"expense_id"`
}

type FetchExpenseOutput struct {
	Expenses []*expensespb.Expense `json:"expenses"`
}

func (this Api) CreateExpense(w http.ResponseWriter, r *http.Request) {
	var expense Expense
	json.NewDecoder(r.Body).Decode(&expense)

	worker := this.LoadBalancer.Next()
	if worker == nil {
		this.responseError(w, 400, fmt.Errorf("No Worker Available"))
		return
	}

	req := &expensespb.CreateExpenseRequest{
		Name:          expense.Name,
		Amount:        expense.Amount,
		BankAccountId: expense.BankAccountId,
	}

	res, err := worker.ExpenseService.CreateExpense(context.Background(), req)
	if err != nil {
		this.responseError(w, 400, err)
		return
	}

	responseJson(w, ExpenseOutput{ExpenseId: res.ExpenseId})
}

func (this Api) FetchExpense(w http.ResponseWriter, r *http.Request) {
	var expense Expense
	json.NewDecoder(r.Body).Decode(&expense)

	worker := this.LoadBalancer.Next()
	if worker == nil {
		this.responseError(w, 400, fmt.Errorf("No Worker Available"))
		return
	}

	req := &expensespb.FetchExpensesRequest{
		BankAccountId: expense.BankAccountId,
	}

	res, err := worker.ExpenseService.FetchExpenses(context.Background(), req)
	if err != nil {
		this.responseError(w, 400, err)
		return
	}

	responseJson(w, FetchExpenseOutput{Expenses: res.Expenses})
}

func (this Api) DeleteExpense(w http.ResponseWriter, r *http.Request) {
	var expense Expense
	json.NewDecoder(r.Body).Decode(&expense)

	worker := this.LoadBalancer.Next()
	if worker == nil {
		this.responseError(w, 400, fmt.Errorf("No Worker Available"))
		return
	}

	req := &expensespb.DeleteExpenseRequest{
		Name:          expense.Name,
		BankAccountId: expense.BankAccountId,
	}

	_, err := worker.ExpenseService.DeleteExpense(context.Background(), req)
	if err != nil {
		this.responseError(w, 400, err)
		return
	}

	responseJson(w, ApiSuccess{Status: "success"})
}
