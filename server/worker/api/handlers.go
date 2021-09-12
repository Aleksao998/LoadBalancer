package api

import "database/sql"

type Api struct {
	Database *sql.DB
}
