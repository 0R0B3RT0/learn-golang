package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	stringConexao := "golang:golang@/devbook?charset=utf8&parseTime=True&loc=Local"
	db, error := sql.Open("mysql", stringConexao)

	if error != nil {
		log.Fatal(error)
	}

	if error = db.Ping(); error != nil {
		log.Fatal(error)
	}

	defer db.Close()

	fmt.Println("Conexão aberta!")
}
