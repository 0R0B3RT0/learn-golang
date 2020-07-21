package main

import (
	"fmt"
)

func main() {
	i := 0
loop:
	for ; i < 10; i++ {
		fmt.Printf("for i = %d\n", i)

		switch i {
		case 2, 3:
			fmt.Printf("Break switch i = %d\n", i)
			break
		case 5:
			fmt.Printf("Break for i=%d\n", i)
			break loop
		}
	}
}
