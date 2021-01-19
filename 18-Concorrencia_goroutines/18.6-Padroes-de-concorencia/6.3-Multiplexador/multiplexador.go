package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Generators")

	canal1 := escrever("Fala galera do canal 1 !!!")
	canal2 := escrever("E aew povinho do canal 2 !!!")

	canalUnico := multiplexador(canal1, canal2)

	for i := 0; i < 10; i++ {
		fmt.Println(<-canalUnico)
	}

}

func multiplexador(canal1, canal2 <-chan string) chan string {
	retorno := make(chan string)

	go func() {
		for {
			select {
			case valor := <-canal1:
				retorno <- valor
			case valor := <-canal2:
				retorno <- valor
			}
		}
	}()

	return retorno
}

func escrever(texto string) chan string {
	canal := make(chan string)
	go func() {
		for {
			canal <- fmt.Sprintf("O valor recebido foi: %s", texto)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(2000)))
		}
	}()

	return canal
}
