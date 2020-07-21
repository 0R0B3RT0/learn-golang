package main

import (
	"fmt"
)

func main() {
	var slice []int
	numbers := []int{1, 2, 3, 4, 5}
	primeNumbers := []int{2, 3, 5, 7, 11, 13}
	names := []string{}

	fmt.Println("Empty slice: ", slice)
	fmt.Println("Filled slice: ", numbers, primeNumbers)
	fmt.Println("Empty slice of string: ", names)

	slice = make([]int, 10)
	fmt.Println("Make slice with 10 values: ", slice, len(slice), cap(slice))
	slice = make([]int, 10, 20)
	fmt.Println("Make slice with 10 values: ", slice, len(slice), cap(slice))
}
