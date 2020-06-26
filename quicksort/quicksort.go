package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	input := os.Args[1:]
	numbers := make([]int, len(input))

	for i, n := range input {
		number, err := strconv.Atoi(n)
		if err != nil {
			fmt.Printf("'%s' isn't valid number\n", n)
			os.Exit(1)
		}
		numbers[i] = number
	}

	fmt.Println(quicksort(numbers))
}

func quicksort(numbers []int) []int {
	if len(numbers) <= 1 {
		return numbers
	}

	values := make([]int, len(numbers))
	copy(values, numbers)

	pivotIndex := len(values) / 2
	pivot := values[pivotIndex]
	values = append(values[:pivotIndex], values[pivotIndex+1:]...)
	minors, biggers := partition(values, pivot)

	return append(append(quicksort(minors), pivot), quicksort(biggers)...)
}

func partition(values []int, pivot int) (minors []int, biggers []int) {
	for _, v := range values {
		if pivot >= v {
			minors = append(minors, v)
		} else {
			biggers = append(biggers, v)
		}
	}

	return minors, biggers
}
