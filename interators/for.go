package main

import (
	"fmt"
)

func main() {
	a, b := 0, 10
	fmt.Println("Values\nA: ", a, " B: ", b)

	examploFor1(a, b)
	examploFor2(a, b)
	examploFor3(a, b)

	slice := []int{3, 2, 1}
	fmt.Println("\n\nValues: ", slice)
	examploForWithRange(slice)
	examploForWithRangeOnlyIndex(slice)
	examploForWithRangeOnlyValue(slice)
	examploForWithRangeOnlyInterate(slice)
}

func examploFor1(a, b int) {
	for a < b {
		a += 1
	}
	fmt.Println("\nfor a < b")
	fmt.Println("A: ", a, " B: ", b)
}

func examploFor2(a, b int) {
	for i := 0; i < b; i++ {
	}
	fmt.Println("\nfor i := 0; i < b; i++")
	fmt.Println("A: ", a, " B: ", b)
}

func examploFor3(a, b int) {
	for ; a < b; a++ {
	}
	fmt.Println("\nfor ; a < b; a++")
	fmt.Println("A: ", a, " B: ", b)
}

func examploForWithRange(numbers []int) {
	fmt.Println("\nfor index, value := range numbers")
	for index, value := range numbers {
		fmt.Println("Index: ", index, "Value: ", value)
	}
}

func examploForWithRangeOnlyIndex(numbers []int) {
	fmt.Println("\nfor index, _ := range numbers")
	for index, _ := range numbers {
		fmt.Println("Index: ", index)
	}
}

func examploForWithRangeOnlyValue(numbers []int) {
	fmt.Println("\nfor _, value := range numbers")
	for _, value := range numbers {
		fmt.Println("Value: ", value)
	}
}

func examploForWithRangeOnlyInterate(numbers []int) {
	fmt.Println("\nfor range numbers")
	for range numbers {
		fmt.Println("Only interate")
	}

	fmt.Println("\nfor _ = range numbers")
	for _ = range numbers {
		fmt.Println("Only interate")
	}
}
