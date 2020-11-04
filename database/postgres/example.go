package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "password"
	dbname   = "inventory"
)

type mov struct {
	ID string
}

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	var myMov mov
	sql := "select id from movement"
	err = db.QueryRow(sql).Scan(&myMov.ID)
	if err != nil {
		fmt.Println("Failed to execute query: ", err)
	}

	fmt.Println(myMov)

	fmt.Println("Successfully connected!")
}
