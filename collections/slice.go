package main

import (
	"fmt"
)

func main() {
	var slice []int
	numbers := []int{1, 2, 3, 4, 6}
	primeNumbers := []int{2, 3, 5, 7, 11, 13}
	names := []string{}

	fmt.Println("Empty slice: ", slice)
	fmt.Println("Filled slice: ", numbers, primeNumbers)
	fmt.Println("Empty slice of string: ", names)

	slice = make([]int, 10)
	fmt.Println("\nmake([]int, 10) = ", slice, len(slice), cap(slice))
	slice = make([]int, 10, 20)
	fmt.Println("make([]int, 10, 20) = ", slice, len(slice), cap(slice))

	fmt.Println("\nSlice slice[start : end]")
	fmt.Println("numbers = ", numbers)
	fmt.Println("numbers[:2] = ", numbers[:2])
	fmt.Println("numbers[1:2] = ", numbers[1:2])
	fmt.Println("numbers[1:] = ", numbers[1:])
	fmt.Println("numbers[1:len(numbers)-1] = ", numbers[1:len(numbers)-1])

	fmt.Println("\nAppend append(slice, element ...)")
	fmt.Println("numbers = ", numbers)
	fmt.Println("append(numbers, 7, 8, 9) = ", append(numbers, 7, 8, 9))
	fmt.Println("append(numbers, []int{7, 8, 9}...) = ", append(numbers, []int{7, 8, 9}...))
	fmt.Println("append([]int{0}, numbers...) = ", append([]int{0}, numbers...))
	fmt.Println("append(numbers[:4], append([]int{5}, numbers[4:]...)...) = ", append(numbers[:4], append([]int{5}, numbers[4:]...)...))

	fmt.Println("\nCopy copy(destiny, origin) int")
	fmt.Println("numbers = ", numbers)
	newSlice := make([]int, len(numbers))
	fmt.Println("New slice before copy = ", newSlice)
	copy(newSlice, numbers)
	fmt.Println("New slice after copy: ", newSlice)
}
