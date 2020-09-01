package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
)

const (
	b3StockNegotiationURL = "https://arquivos.b3.com.br/apinegocios/ticker/%s/%s"
)

func main() {
	s, err := stockQuote("ITUB3", "2020-07-14")

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Ação: %s\tCotação: %.2f\tData: %s\n\n", s.name, s.price, s.date)
}

func stockQuote(name, date string) (stock, error) {
	body, err := DoGet(makeURL(b3StockNegotiationURL, name, date))

	requestStruct := parse(body)

	r, err := requestStruct.getLastCotation()

	var s stock
	if err != nil {
		return s, err
	}

	quoteDate := fmt.Sprintf("%s %s", r.getDate(), r.getTime())
	s = stock{r.getStock(), r.getPrice(), quoteDate}
	return s, nil
}

func makeURL(baseURL, stock, day string) string {
	return fmt.Sprintf(baseURL, stock, day)
}

func (requestStruct requestStruct) getLastCotation() (values, error) {
	if len(requestStruct.values) > 0 {
		return requestStruct.values[0], nil
	}
	return nil, errors.New("There is not values")
}

func parse(body []byte) requestStruct {
	var requestStruct requestStruct
	json.Unmarshal(body, &requestStruct)
	return requestStruct
}

type stock struct {
	name  string
	price float64
	date  string
}

type requestStruct struct {
	Name         string //ITUB3
	FriendlyName string //ITUB3
	Columns      []column
	values       []values
}

type column struct {
	Name            string //TickerSymbol
	FriendlyName    string //Papel
	FriendlyNamePt  string //Papel
	FriendlyNameEn  string //Symbol
	Type            string //5
	ColumnAlignment string //1
	ValueAlignment  string //2
}

type values []interface{}

func (v values) getStock() string {
	return fmt.Sprintf("%v", v[0])
}

func (v values) getPrice() float64 {
	str := fmt.Sprintf("%v", v[2])
	value, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return 0
	}
	return value
}

func (v values) getDate() string {
	return fmt.Sprintf("%v", v[4])
}

func (v values) getTime() string {
	return fmt.Sprintf("%v", v[5])
}
