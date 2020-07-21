package main

import (
	"errors"
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

	printRemovedValue(stack.UnStack())
	printRemovedValue(stack.UnStack())
	printRemovedValue(stack.UnStack())
	printRemovedValue(stack.UnStack())
	printRemovedValue(stack.UnStack())
	printRemovedValue(stack.UnStack())

	stack.PrintInformations()
}

func printRemovedValue(value interface{}, err error) {
	if err == nil {
		fmt.Println("Removed value: ", value)
	} else {
		fmt.Println("Occur error: ", err)
	}
}

type Stack struct {
	values []interface{}
}

func (stack Stack) Tamanho() int {
	return len(stack.values)
}

func (stack Stack) Empty() bool {
	return stack.Tamanho() == 0
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

func (stack *Stack) UnStack() (interface{}, error) {
	if stack.Empty() {
		return nil, errors.New("Empty stack!")
	}
	value := stack.values[len(stack.values)-1]
	stack.values = stack.values[:len(stack.values)-1]

	return value, nil
}
