package main

import "fmt"

func calculos(v1, v2 int) (soma, subtracao int) {
	soma = v1 + v2
	subtracao = v1 - v2

	return
}

func main() {
	fmt.Println(calculos(12, 6))
}
