package main

import "fmt"

func main() {
	canal := make(chan string, 2)

	canal <- "Texto 1"
	canal <- "Texto 2"
	// close(canal)
	// canal <- "Texto 3"

	fmt.Println(<-canal)
	fmt.Println(<-canal)

	canal <- "Texto 3"
	canal <- "Texto 4"

	fmt.Println(<-canal)
	fmt.Println(<-canal)
}
