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

	sql := "select count(1) from movement"
	rows, err := db.Query(sql)
	if err != nil {
		fmt.Println("Failed to execute query: ", err)
	}
	rows.Next()
	var count int
	rows.Scan(&count)
	fmt.Println(count)

	fmt.Println("Successfully connected!")
}
