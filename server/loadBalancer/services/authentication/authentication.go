package authentication

import (
	"database/sql"
	"fmt"

	"github.com/Aleksao998/LoadBalancer/config"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        int64
	FirstName string
	LastName  string
	Email     string
	Password  string
}

func DoesUserExists(email string, db *sql.DB) (bool, error) {
	var count int
	row := db.QueryRow("select count(*) from users where email = $1", email)
	err := row.Scan(&count)
	if err != nil {
		return false, fmt.Errorf("Sql Error", err)
	}
	if count != 0 {
		return false, nil
	}
	return true, nil
}

func CreateHashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", fmt.Errorf("Bcrypt Error", err)
	}

	return string(hash), nil
}

func RegisterUser(user User, db *sql.DB) error {
	var newUser User
	stm := "insert into users (email, password, first_name, last_name) values($1, $2, $3, $4) returning user_id"
	err := db.QueryRow(stm, user.Email, user.Password, user.FirstName, user.LastName).Scan(&newUser.ID)
	if err != nil && err != sql.ErrNoRows {
		return fmt.Errorf("SQL Error", err)
	}
	return nil
}

func CompareHashAndPassword(hashedPassword string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return false
	}
	return true
}

func GenerateToken(email string) (string, error) {
	secret := config.Config.JWT.Secret

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
	})

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", fmt.Errorf("JWT Error", err)
	}

	return tokenString, nil
}

func GetUserData(email string, db *sql.DB) (string, error) {
	var password string

	row := db.QueryRow("select password from users where email = $1", email)
	err := row.Scan(&password)

	if err == sql.ErrNoRows {
		return "", nil
	}
	if err != nil {
		return "", fmt.Errorf("SQL Error", err)
	}
	return password, nil
}
