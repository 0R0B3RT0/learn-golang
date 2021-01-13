package main

import (
	"fmt"
	"time"
)

func main() {
	go escrever("Um texto") //goroutine
	escrever("Outro texto")
}

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
