package main

import (
	"fmt"
)

func main() {
	var i int
externo:
	for {
		for i = 0; i < 10; i++ {
			if i == 5 {
				break externo
			}
		}
	}
	fmt.Println(i)
}
