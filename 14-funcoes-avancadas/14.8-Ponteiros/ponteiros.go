package main

import "fmt"

func inverteSinal(numero int) int {
	fmt.Println("Inverte sem ponteiro")
	numero *= -1
	return numero
}

func inverteSinalComPonteiro(numero *int) {
	fmt.Println("Inverte com ponteiro")
	*numero = *numero * -1
}

func main() {
	fmt.Println("Ponteiros")
	numero := 10
	fmt.Println(inverteSinal(numero))
	fmt.Println(numero)

	inverteSinalComPonteiro(&numero)
	fmt.Println(numero)
}
