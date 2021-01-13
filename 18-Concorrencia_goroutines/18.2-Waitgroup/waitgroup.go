package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var waitgroup sync.WaitGroup
	waitgroup.Add(2)

	go func() {
		escrever("Um texto")
		waitgroup.Done()
	}()

	go func() {
		escrever("Outro texto")
		waitgroup.Done()
	}()

	waitgroup.Wait()
}

func escrever(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
