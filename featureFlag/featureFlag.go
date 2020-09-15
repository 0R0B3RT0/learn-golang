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
	password = "postgres"
	dbname   = "feature-flag"
)

type feature struct {
	ID          string
	Value       bool
	ExternalID  string
	FeatureName string
}

func main() {

	myFeat, err := findFeature("a35909d9-d20e-4dad-950b-8269184d8f35", "funcionalidade um")
	if err != nil {
		fmt.Println("Failed to execute query: ", err)
	}

	fmt.Println(myFeat)

	fmt.Println("Successfully connected!")
}

func findFeature(externalID, featureName string) (feature, error) {
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

	var myFeat feature
	sql := "select id, external_id, feature_name, value from feature"
	err = db.QueryRow(sql).Scan(&myFeat.ID, &myFeat.ExternalID, &myFeat.FeatureName, &myFeat.Value)

	return myFeat, err
}
