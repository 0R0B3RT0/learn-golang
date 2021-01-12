package main

import "fmt"

func main() {
	var v1 int = 10
	var v2 int = v1

	fmt.Println(v1, v2)

	v2++
	fmt.Println(v1, v2)

	var v3 int
	var ponteiro *int

	fmt.Println(v3, ponteiro)

	v3 = 100
	ponteiro = &v3

	// Valor do ponteiro
	fmt.Println(v3, *ponteiro)

	// Endereço do ponteiro
	fmt.Println(v3, ponteiro)

	v3 = 200

	fmt.Println(v3, ponteiro)
	fmt.Println(v3, *ponteiro)

}
