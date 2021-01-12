package main

import "fmt"

var text string

func init() {
	fmt.Println("Executando a função Init")
	text = "String iniciada"
}
func main() {
	fmt.Println("Executando a função main")
	fmt.Println(text)
}
