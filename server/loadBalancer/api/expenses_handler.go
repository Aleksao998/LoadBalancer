package api

import (
	"fmt"
	"net/http"
)

func (this Api) CreateExpense(w http.ResponseWriter, r *http.Request) {
	//specify status code
	w.WriteHeader(http.StatusOK)

	//update response writer
	fmt.Fprintf(w, "Api for CreateExpense")
}

func (this Api) FetchExpense(w http.ResponseWriter, r *http.Request) {
	//specify status code
	w.WriteHeader(http.StatusOK)

	//update response writer
	fmt.Fprintf(w, "Api for FetchExpense")
}

func (this Api) DeleteExpense(w http.ResponseWriter, r *http.Request) {
	//specify status code
	w.WriteHeader(http.StatusOK)

	//update response writer
	fmt.Fprintf(w, "Api for DeleteExpense")
}
