package main

import "fmt"

func main() {
	fmt.Println(somaESubtrai(10, 15))

	res1, res2 := somaESubtrai(1, 2)
	fmt.Println(res1, res2)

	fmt.Println(calculos(10, 15))

	soma, _, _, _ := calculos(10, 15)
	fmt.Println("soma: ", soma)

	_, subtrai, _, _ := calculos(10, 15)
	fmt.Println("subtrai: ", subtrai)

	_, _, multiplica, _ := calculos(10, 15)
	fmt.Println("multiplica: ", multiplica)

	_, _, _, divide := calculos(10, 15)
	fmt.Println("divide: ", divide)
}

func somaESubtrai(v1, v2 int8) (int8, int8) {
	return (v1 + v2), (v1 - v2)
}

func calculos(v1, v2 int) (int, int, int, int) {
	return (v1 + v2), (v1 - v2), (v1 * v2), (v1 / v2)
}
