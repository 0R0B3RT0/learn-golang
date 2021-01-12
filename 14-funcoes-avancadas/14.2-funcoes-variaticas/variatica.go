package main

import "fmt"

func soma(valores ...int) (resultado int) {
	for _, v := range valores {
		resultado += v
	}
	return
}

func main() {
	fmt.Println(soma(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
	fmt.Println(soma())
}
