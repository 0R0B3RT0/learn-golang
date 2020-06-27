package main

import (
	// "errors"
	"fmt"
)

func main() {
	stack := Stack{}

	stack.PrintInformations()

	stack.StackUp("Learning")
	stack.StackUp("GoLang")
	stack.StackUp("in")
	stack.StackUp("2020")
	stack.StackUp("!!!")

	stack.PrintInformations()

	stack.UnStack()
	stack.UnStack()
	stack.UnStack()
	stack.UnStack()
	stack.UnStack()

	stack.PrintInformations()
}

type Stack struct {
	values []interface{}
}

func (stack Stack) Tamanho() int {
	return len(stack.values)
}

func (stack Stack) PrintInformations() {
	if stack.Tamanho() == 0 {
		fmt.Println("Empty stack!!!")
	} else {
		fmt.Println("Stack created with length ", stack.Tamanho())
	}
}

func (stack *Stack) StackUp(value interface{}) {
	stack.values = append(stack.values, value)
}

func (stack *Stack) UnStack() {
	value := stack.values[len(stack.values)-1]
	stack.values = stack.values[:len(stack.values)-1]

	fmt.Println("Removed value: ", value)
}
