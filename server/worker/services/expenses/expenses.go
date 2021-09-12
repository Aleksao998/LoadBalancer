package expenses

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/Aleksao998/LoadBalancer/worker/api/pb/expensespb"
)

func DoesExpenseExists(expenseName string, bankAccountId string, db *sql.DB) (bool, error) {
	var count int

	id, err := strconv.Atoi(bankAccountId)
	if err != nil {
		return false, fmt.Errorf("Argment Error", err)
	}

	row := db.QueryRow("select count(*) from expenses where name = $1 and bank_account_id = $2", expenseName, id)
	err = row.Scan(&count)
	if err != nil {
		return false, fmt.Errorf("Sql Error", err)
	}
	if count != 0 {
		return false, nil
	}
	return true, nil
}

func RegisterExpense(expenseName string, bankAccountId string, amount float32, db *sql.DB) (int, error) {
	var expense_id int
	id, err := strconv.Atoi(bankAccountId)
	if err != nil {
		return -1, fmt.Errorf("Argument Error", err)
	}
	stm := "insert into expenses (name, amount, bank_account_id) values($1, $2, $3) returning expense_id"
	err = db.QueryRow(stm, expenseName, amount, id).Scan(&expense_id)
	if err != nil && err != sql.ErrNoRows {
		return -1, fmt.Errorf("SQL Error", err)
	}
	return expense_id, nil
}

func FetchExpenses(bankAccountId string, db *sql.DB) ([]*expensespb.Expense, error) {

	var expenses []*expensespb.Expense

	id, err := strconv.Atoi(bankAccountId)
	if err != nil {
		return nil, fmt.Errorf("Argument Error", err)
	}
	rows, err := db.Query("select name, amount from expenses where bank_account_id = $1", id)
	if err != nil {
		return nil, fmt.Errorf("Sql Error", err)
	}
	type expenseSql struct {
		name   string
		amount float32
	}

	for rows.Next() {
		var expenseSqlData expenseSql
		var expense expensespb.Expense
		rows.Scan(&expenseSqlData.name, &expenseSqlData.amount)
		if err != nil {
			return nil, fmt.Errorf("Sql Error", err)
		}
		fmt.Println(expenseSqlData)
		expense.Amount = expenseSqlData.amount
		expense.Name = expenseSqlData.name
		expenses = append(expenses, &expense)
	}

	return expenses, nil
}

func DeleteBankAccount(expenseName string, bankAccountId string, db *sql.DB) error {
	id, err := strconv.Atoi(bankAccountId)
	if err != nil {
		return fmt.Errorf("Argument Error", err)
	}
	sqlDelete := "delete from expenses where name = $1 and bank_account_id = $2"
	res, err := db.Exec(sqlDelete, expenseName, id)
	if err != nil {
		return fmt.Errorf("SQL Error", err)
	}
	count, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("SQL Error", err)
	}
	if count == 0 {
		return fmt.Errorf("No rows found")
	}
	return nil
}
