package banco

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func Conect() (*sql.DB, error) {
	connectionString := "golang:golang@/devbook"

	db, error := sql.Open("mysql", connectionString)
	if error != nil {
		return nil, error
	}

	if error = db.Ping(); error != nil {
		return nil, error
	}

	return db, nil
}
