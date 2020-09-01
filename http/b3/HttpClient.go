package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// DoGet execute an http GET with the informed URL and return the []byte or error
func DoGet(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Request with fails, error: ", err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Fail to read body, error: ", err)
		return nil, err
	}

	return body, nil
}
