package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 3{
		fmt.Println("Use: convert <values> <unit>")
		os.Exit(1)
	}
	var destinationUnit string
	originUnit := os.Args[len(os.Args)-1]
	originValues := os.Args[1 : len(os.Args)-1]

	if originUnit == "celsios"{
		destinationUnit = "fahrenheit"
	}else if originUnit == "quilometers"{
		destinationUnit = "miles"
	} else {
		fmt.Printf("%s is invalid!", originUnit)
		os.Exit(1)
	}
	
	for i, v := range originValues {
		originValue, err := strconv.ParseFloat(v, 64)

		if(err != nil){
			fmt.Printf("The value %s at position %d it's invalid number!\n", v, i)
			os.Exit(1)
		}

		var destinationValue float64
		if originUnit == "celsios"{
			destinationValue = originValue * 1.8 + 32
		}else {
			destinationValue = originValue / 1.60934
		}

		fmt.Printf("%.2f %s = %.2f %s\n", originValue, originUnit, destinationValue, destinationUnit)
	}
}