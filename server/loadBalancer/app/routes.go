package app

import (
	"fmt"
	"net/http"

	"github.com/Aleksao998/LoadBalancer/api"
	"github.com/Aleksao998/LoadBalancer/config"
	"github.com/gorilla/mux"
)

func (this App) starHttpServer(api *api.Api) {
	router := mux.NewRouter()

	this.registerHttpRoutes(router, *api)

	server := &http.Server{Addr: ":" + config.Config.Web.Port, Handler: router}
	if err := server.ListenAndServe(); err != nil {
		panic(err.Error())
	}
	fmt.Printf("HTTP Server Started on port: " + config.Config.Web.Port + "\n")
}

func (this App) registerHttpRoutes(router *mux.Router, api api.Api) {

	router.HandleFunc("/create-bank-account", api.TokenVerifyMiddleWare(api.CreateBankAccount)).Methods("POST")
	router.HandleFunc("/delete-bank-account", api.TokenVerifyMiddleWare(api.DeleteBankAccount)).Methods("POST")
	router.HandleFunc("/fetch-bank-account", api.TokenVerifyMiddleWare(api.FetchBankAccount)).Methods("GET")

	router.HandleFunc("/create-expense", api.TokenVerifyMiddleWare(api.CreateExpense)).Methods("POST")
	router.HandleFunc("/delete-expense", api.TokenVerifyMiddleWare(api.DeleteExpense)).Methods("POST")
	router.HandleFunc("/fetch-expenses", api.TokenVerifyMiddleWare(api.FetchExpense)).Methods("GET")

	router.HandleFunc("/login", api.Login).Methods("POST")
	router.HandleFunc("/register", api.Register).Methods("POST")

	http.Handle("/", router)
}
