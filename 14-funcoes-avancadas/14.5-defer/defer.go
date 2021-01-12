package main

import "fmt"

func func1() {
	fmt.Println("Função 1")
}

func func2() {
	fmt.Println("Função 2")
}

func func3() {
	fmt.Println("Função 3")
}

func main() {
	defer func1()
	defer func2()
	func3()
}
