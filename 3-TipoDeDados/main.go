package main

import (
	"errors"
	"fmt"
)

func main() {
	var numero1 int64 = 100000000000000000
	fmt.Println(numero1)

	var numero2 int32 = 1000000000
	fmt.Println(numero2)

	var numero3 rune = 1000000000
	fmt.Println(numero3)

	var numero4 byte = 100
	fmt.Println(numero4)

	var numero5 uint8 = 100
	fmt.Println(numero5)

	var numeroReal float32 = 172376.98
	fmt.Println(numeroReal)

	numeroReal2 := 726.98
	fmt.Println(numeroReal2)

	var str string = "Texto"
	fmt.Println(str)

	str2 := "Texto 2"
	fmt.Println(str2)

	char := 'k'
	fmt.Println(char)

	char2 := '\''
	fmt.Println(char2)

	var boolean bool
	fmt.Println(boolean)

	var erro error
	fmt.Println(erro)

	var erro2 error = errors.New("Internal error")
	fmt.Println(erro2)
}
