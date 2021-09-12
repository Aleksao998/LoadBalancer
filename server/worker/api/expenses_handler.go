package api

import (
	"fmt"
	"strconv"
	"context"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/codes"
	"github.com/Aleksao998/LoadBalancer/worker/services/expenses"
	"github.com/Aleksao998/LoadBalancer/worker/api/pb/expensespb"
)

func (this *Api) CreateExpense(ctx context.Context, req *expensespb.CreateExpenseRequest) (*expensespb.CreateExpenseResponse, error) {
	name := req.GetName()
	bankAccountId := req.GetBankAccountId()
	amount := req.GetAmount()
	var result string

	ok, err := expenses.DoesExpenseExists(name, bankAccountId, this.Database)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal, err.Error(),
		)

	}
	if !ok {
		return nil, status.Errorf(
			codes.AlreadyExists, fmt.Sprintf("Expense already exists"),
		)
	}

	expense_id, err := expenses.RegisterExpense(name,bankAccountId,amount, this.Database)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,err.Error(),
		)
	}

	result = strconv.Itoa(expense_id)
	res := expensespb.CreateExpenseResponse{
		ExpenseId: result,
	}

	return &res, nil
}

func (this *Api) DeleteExpense(ctx context.Context, req *expensespb.DeleteExpenseRequest) (*expensespb.DeleteExpensetResponse, error) {
	name := req.GetName()
	bankAccountId := req.GetBankAccountId()

	err := expenses.DeleteBankAccount(name,bankAccountId, this.Database)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,err.Error(),
		)
	}

	res := expensespb.DeleteExpensetResponse{
	}

	return &res, nil
}


func (this *Api) FetchExpenses(ctx context.Context, req *expensespb.FetchExpensesRequest) (*expensespb.FetchExpensesResponse, error) {
	bankAccountId := req.GetBankAccountId()

	expenses, err := expenses.FetchExpenses(bankAccountId, this.Database)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,err.Error(),
		)
	}

	res := expensespb.FetchExpensesResponse{
		Expenses: expenses,
	}

	return &res, nil
}
