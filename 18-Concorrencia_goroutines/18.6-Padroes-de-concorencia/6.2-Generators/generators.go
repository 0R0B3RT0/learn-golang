package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Generators")

	canal := escrever("Fala galera!!!")

	for i := 0; i < 10; i++ {
		fmt.Println(<-canal)
	}

}

func escrever(texto string) chan string {
	canal := make(chan string)
	go func() {
		for {
			canal <- fmt.Sprintf("O valor recebido foi: %s", texto)
			time.Sleep(time.Millisecond * 500)
		}
	}()

	return canal
}
