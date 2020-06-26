package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 3{
		fmt.Println("Use: convert <originUnit> <destionationUnit> <values>")
		os.Exit(1)
	}

	originUnit := os.Args[1]
	destinationUnit := os.Args[2]
	originValues := os.Args[3 : len(os.Args)]

	if originUnit != "celsius" && originUnit != "fahrenheit" && originUnit != "quilometers" && originUnit != "miles"{
		fmt.Printf("%s is invalid!", originUnit)
		os.Exit(1)
	}
	if destinationUnit != "celsius" && destinationUnit != "fahrenheit" && destinationUnit != "quilometers" && destinationUnit != "miles"{
		fmt.Printf("%s is invalid!", destinationUnit)
		os.Exit(1)
	}
	
	for i, v := range originValues {
		originValue, err := strconv.ParseFloat(v, 64)

		if(err != nil){
			fmt.Printf("The value %s at position %d it's invalid number!\n", v, i)
			os.Exit(1)
		}

		destinationValue := convert(originValue, originUnit, destinationUnit)

		fmt.Printf("%.2f %s = %.2f %s\n", originValue, originUnit, destinationValue, destinationUnit)
	}
}

func convert(value float64, origin, destination string) float64{
	var destinationValue float64

	if origin == "celsius" && destination == "fahrenheit"{
		destinationValue = value * 1.8 + 32
	} else if origin == "fahrenheit" && destination == "celsius"{
		destinationValue = (value - 32) / 1.8
	}else if origin == "quilometers" && destination == "miles"{
		destinationValue = value / 1.60934
	} else if origin == "miles" && destination == "quilometers"{
		destinationValue = value * 1.60934
	} else {
		fmt.Printf("Does not possible convert %s to %s\n", origin, destination)
		os.Exit(1)
	}

	return destinationValue
}