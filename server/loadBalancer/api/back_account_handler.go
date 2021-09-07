package api

import (
	"fmt"
	"net/http"
)

func (this Api) CreateBankAccount(w http.ResponseWriter, r *http.Request) {
	//specify status code
	w.WriteHeader(http.StatusOK)

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
