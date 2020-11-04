package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"net/url"
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
	http.HandleFunc("/", apiResponse)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func apiResponse(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	query, _ := url.ParseQuery(r.URL.RawQuery)

	log.Println(query["feature"][0])
	path := r.URL.RequestURI()
	log.Println("before parse: ", path)
	parsedPath, _ := url.ParseRequestURI(path)
	log.Println(" after parse: ", parsedPath)

	switch r.Method {
	case "GET":
		featureName := query["feature"][0]
		myFeat, err := findFeature("a35909d9-d20e-4dad-950b-8269184d8f35", featureName)
		if err != nil {
			log.Println("Failed to execute query: ", err)
		}
		log.Println(myFeat)

		myFeatJSON, err := json.Marshal(myFeat)
		if err != nil {
			log.Println("Failed to json marshal: ", err)
		}
		w.Write([]byte(myFeatJSON))
	default:
		w.Write([]byte("Method not supported"))
	}

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
