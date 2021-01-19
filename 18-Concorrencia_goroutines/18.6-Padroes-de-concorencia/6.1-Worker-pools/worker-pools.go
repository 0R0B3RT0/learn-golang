package main

import "fmt"

func main() {
	total := 45
	tarefas, resultados := make(chan int, total), make(chan int, total)

	go worker(tarefas, resultados)
	go worker(tarefas, resultados)

	for i := 0; i < total; i++ {
		tarefas <- i
	}

	close(tarefas)

	for i := 0; i < total; i++ {
		fmt.Println(<-resultados)
	}
}

func worker(tarefas <-chan int, resultados chan<- int) {
	for numero := range tarefas {
		resultados <- fibonacci(numero)
	}
}

func fibonacci(posicao int) int {
	if posicao <= 1 {
		return posicao
	}

	return fibonacci(posicao-2) + fibonacci(posicao-1)
}
