package app

import (
	"database/sql"
	"fmt"

	"github.com/Aleksao998/LoadBalancer/worker/config"
	_ "github.com/lib/pq"
)

func (this App) getConnection() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config.Config.Database.Host, config.Config.Database.Port, config.Config.Database.User, config.Config.Database.Pwd, config.Config.Database.Db)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return db
}
