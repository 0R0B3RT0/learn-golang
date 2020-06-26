package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	input := os.Args[1:]

	statistics := getStatistics(input)

	printStatistics(statistics)
}

func getStatistics(words []string) map[string]int {
	statistics := make(map[string]int)

	for _, word := range words {
		word = strings.ToUpper(string(word))
		for _, letter := range word {
			letter := string(letter)
			counter, exist := statistics[letter]
			if exist {
				statistics[letter] = counter + 1
			} else {
				statistics[letter] = 1
			}
		}
	}

	return statistics
}

func printStatistics(statistics map[string]int) {
	for key, value := range statistics {
		fmt.Printf("%s = %d\n", key, value)
	}
}
