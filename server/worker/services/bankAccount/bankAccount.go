package bankAccount

import (
	"database/sql"
	"fmt"
	"strconv"
)

func DoesBankAccountExists(bankAccountName string, userId string ,db *sql.DB) (bool, error) {
	var count int

	id, err := strconv.Atoi(userId); 
	if err != nil {
		return false, fmt.Errorf("Argment Error", err)
	}

	row := db.QueryRow("select count(*) from bank_account where name = $1 and user_id = $2", bankAccountName, id)
	err = row.Scan(&count)
	if err != nil {
		return false, fmt.Errorf("Sql Error", err)
	}
	if count == 0 {
		return false, nil
	}
	return true, nil
}

func RegisterBankAccount(bankAccountName string, userId string, db *sql.DB) (int,error) {
	var user_id int
	id, err := strconv.Atoi(userId); 
	if err != nil {
		return -1, fmt.Errorf("Argument Error", err)
	}
	stm := "insert into bank_account (name, user_id) values($1, $2) returning bank_account_id"
	err = db.QueryRow(stm, bankAccountName, id).Scan(&user_id)
	if err != nil && err != sql.ErrNoRows {
		return -1,fmt.Errorf("SQL Error", err)
	}
	return user_id,nil
}

func DeleteBankAccount(bankAccountName string, userId string, db *sql.DB) (error) {
	id, err := strconv.Atoi(userId); 
	if err != nil {
		return fmt.Errorf("Argument Error", err)
	}
	sqlDelete := "delete from bank_account where name = $1 and user_id = $2"
	res, err := db.Exec(sqlDelete,bankAccountName, id)
	if err != nil {
		return fmt.Errorf("SQL Error", err)
	}
	count, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("SQL Error", err)
	}
	if count == 0{
		return fmt.Errorf("No rows found")
	}
	return nil
}



func FetchBankAccountAmount(bankAccountName string, userId string, db *sql.DB) (float32,error) {
	var bankAccountId int
	var totalAmount float32
	id, err := strconv.Atoi(userId); 
	if err != nil {
		return -1, fmt.Errorf("Argment Error", err)
	}

	row := db.QueryRow("select bank_account_id from bank_account where name = $1 and user_id = $2", bankAccountName, id)
	err = row.Scan(&bankAccountId)
	if err != nil {
		return -1, fmt.Errorf("Sql Error", err)
	}

	row = db.QueryRow("select sum(amount) from expenses where bank_account_id = $1",bankAccountId)
	err = row.Scan(&totalAmount)
	if err != nil {
		return -1, fmt.Errorf("Sql Error", err)
	}

	return totalAmount, nil
}