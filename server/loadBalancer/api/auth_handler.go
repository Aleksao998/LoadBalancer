package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	AuthService "github.com/Aleksao998/LoadBalancer/services/authentication"
)

type Status struct {
	Status   int    `json:"status"`
	ErrorMsg string `json:"errorMsg"`
}

type Data struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Token     string `json:"token"`
}

type AuthOutputData struct {
	Status Status `json:"status"`
	Data   Data   `json:"data"`
}

func (this Api) generateAuthOutputData(status int, errorMsg string, firstName string, lastName string, token string) AuthOutputData {
	return AuthOutputData{
		Status: Status{
			Status:   status,
			ErrorMsg: errorMsg,
		},
		Data: Data{
			FirstName: firstName,
			LastName:  lastName,
			Token:     token,
		},
	}
}

func (this Api) Register(w http.ResponseWriter, r *http.Request) {
	var user AuthService.User

	json.NewDecoder(r.Body).Decode(&user)

	if user.Email == "" || user.Password == "" || user.FirstName == "" || user.LastName == "" {
		this.responseError(w, http.StatusBadRequest, fmt.Errorf("Missing filds"))
		return
	}

	ok, err := AuthService.DoesUserExists(user.Email, this.Database)
	if err != nil {
		this.responseError(w, http.StatusInternalServerError, err)
		return
	}
	if !ok {
		this.responseError(w, http.StatusBadRequest, fmt.Errorf("User Already Exists"))
		return
	}

	user.Password, err = AuthService.CreateHashPassword(user.Password)
	if err != nil {
		this.responseError(w, http.StatusInternalServerError, err)
		return
	}

	err = AuthService.RegisterUser(user, this.Database)
	if err != nil {
		this.responseError(w, http.StatusInternalServerError, err)
		return
	}

	token, err := AuthService.GenerateToken(user.Email)
	if err != nil {
		this.responseError(w, http.StatusInternalServerError, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	var outputData = this.generateAuthOutputData(200, "", user.FirstName, user.LastName, token)

	responseJson(w, outputData)
	return

}

func (this Api) Login(w http.ResponseWriter, r *http.Request) {
	var user AuthService.User
	json.NewDecoder(r.Body).Decode(&user)

	if user.Email == "" || user.Password == "" {
		this.responseError(w, http.StatusBadRequest, fmt.Errorf("Missing filds"))
		return
	}

	password := user.Password

	hashedPassword, err := AuthService.GetUserData(user.Email, this.Database)
	if err != nil {
		this.responseError(w, http.StatusInternalServerError, err)
		return
	}

	if hashedPassword == "" {
		this.responseError(w, http.StatusBadRequest, fmt.Errorf("Email does not exists"))
		return
	}

	if !AuthService.CompareHashAndPassword(hashedPassword, password) {
		this.responseError(w, http.StatusBadRequest, fmt.Errorf("Password does not match"))
		return
	}

	token, err := AuthService.GenerateToken(user.Email)
	if err != nil {
		this.responseError(w, http.StatusInternalServerError, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	var outputData = this.generateAuthOutputData(200, "", user.FirstName, user.LastName, token)
	responseJson(w, outputData)
	return
}
