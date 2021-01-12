package main

import "fmt"

func main() {
	func() {
		fmt.Println("Teste função anônima")
	}()

	print := func(texto string) {
		fmt.Println(texto)
	}

	print("Fala galera!!!")
}
