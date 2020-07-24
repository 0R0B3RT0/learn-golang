package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	BaseUrl = "https://arquivos.b3.com.br/apinegocios/ticker/ITUB3/2020-07-24"
)

func main() {
	resp, err := http.Get(BaseUrl)
	if err != nil {
		fmt.Println("Request with fails, error: ", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Fail to read body, error: ", err)
	}

	fmt.Printf("Response: %d\n", resp.StatusCode)
	// fmt.Printf("%s\n", body)

	var requestStruct RequestStruct
	json.Unmarshal(body, &requestStruct)

	// fmt.Printf("Object: %s\n", requestStruct.Values)

	// var last []interface{}
	for _, r := range requestStruct.Values {
		// fmt.Printf("Ação: %s\tCotação: %.2f\tData: %s\tHora: %s\n", r[0], r[2], r[4], r[5])
		fmt.Printf("Ação: %s\tCotação: %.2f\tData: %s\tHora: %s\n", r.getStock(), r.getPrice(), r.getDate(), r.getTime())
	}
}

type RequestStruct struct {
	Name         string //ITUB3
	FriendlyName string //ITUB3
	Columns      []Column
	Values       []Values
}

type Column struct {
	Name            string //TickerSymbol
	FriendlyName    string //Papel
	FriendlyNamePt  string //Papel
	FriendlyNameEn  string //Symbol
	Type            string //5
	ColumnAlignment string //1
	ValueAlignment  string //2
}

type Values []interface{}

func (v Values) getStock() interface{} {
	return v[0]
}

func (v Values) getPrice() interface{} {
	return v[2]
}

func (v Values) getDate() interface{} {
	return v[4]
}

func (v Values) getTime() interface{} {
	return v[5]
}
