package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func testMakeURL(t *testing.T) {
	actual := makeURL("www.teste.com", "stock", "date")
	assert.Equal(t, "www.teste.com/stock/date", actual, "they should be equal")
}

func testGetLastCotation(t *testing.T) {
	v := []values{
		{"stock1", 1, 3.50, 3, "2020-07-26", "12:00:00"},
		{"stock1", 1, 5.50, 3, "2020-07-26", "11:00:00"},
	}
	requestStruct := requestStruct{"name", "friendlyName", nil, v}

	lastNegotiation, err := requestStruct.getLastCotation()

	assert.EqualValues(t, values{"stock1", 1, 3.50, 3, "2020-07-26", "12:00:00"}, lastNegotiation, "they should be equal")
	assert.Nil(t, err, "they should be null")
}
