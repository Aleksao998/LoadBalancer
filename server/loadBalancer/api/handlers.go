package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/Aleksao998/LoadBalancer/config"
	"github.com/Aleksao998/LoadBalancer/services/loadBalancer"
	"github.com/dgrijalva/jwt-go"
)

type Api struct {
	Database     *sql.DB
	LoadBalancer *loadBalancer.WorkersPool
}

type ApiError struct {
	ErrorMsg string `json:"error_message"`
}

type ApiSuccess struct {
	Status string `json:"status"`
}

func responseJson(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		fmt.Printf("ERROR", err)
	}
}

func (this Api) responseError(w http.ResponseWriter, errorCode int, error error) {
	w.WriteHeader(errorCode)
	responseJson(w, ApiError{ErrorMsg: error.Error()})
}

func (this Api) TokenVerifyMiddleWare(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		bearerToken := strings.Split(authHeader, " ")

		if len(bearerToken) == 2 {
			authToken := bearerToken[1]

			token, err := jwt.Parse(authToken, func(token *jwt.Token) (interface{}, error) {

				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("There was an Error")
				}

				return []byte(config.Config.JWT.Secret), nil
			})
			if err != nil {
				this.responseError(w, http.StatusUnauthorized, fmt.Errorf("Invalid Token"))
				return
			}

			if token.Valid {
				next.ServeHTTP(w, r)
			} else {
				this.responseError(w, http.StatusUnauthorized, fmt.Errorf("Invalid Token"))
				return
			}

		} else {
			this.responseError(w, http.StatusUnauthorized, fmt.Errorf("Invalid Token"))
			return
		}
	})
}
