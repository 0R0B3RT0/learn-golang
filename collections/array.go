package main

import (
	"fmt"
)

func main() {
	var array [3]int
	numbers := [5]int{1, 2, 3, 4, 5}
	primeNumbers := [...]int{2, 3, 5, 7, 11, 13}
	names := [2]string{}
	multidimensional := [2][2]int{}
	multidimensionalFilled := [2][2]int{{11, 12}, {21, 22}}

	fmt.Println("Empty array: ", array)
	fmt.Println("Filled slice: ", numbers, primeNumbers)
	fmt.Println("Empty slice of string: ", names)
	fmt.Println("Empty multidimensional array: ", multidimensional)
	fmt.Println("Filled multidimensional array: ", multidimensionalFilled)
}
