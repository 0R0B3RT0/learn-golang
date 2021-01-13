package main

import (
	"fmt"
	"time"
)

func main() {
	canal := make(chan string)

	go escrever("Exemplo de canal simples", canal)

	// Simples
	for {
		message, aberto := <-canal
		if !aberto {
			break
		}
		fmt.Println(message)
	}

	canal = make(chan string)
	go escrever("Exemplo de canal elegante", canal)

	// Elegante
	for message := range canal {
		fmt.Println(message)
	}

	fmt.Println("Finalizado o processamento")
}

func escrever(text string, canal chan string) {
	for i := 0; i < 5; i++ {
		canal <- fmt.Sprintf("%s - count: %d", text, i)
		// fmt.Println(text)
		time.Sleep(time.Second)
	}
	close(canal)
}
