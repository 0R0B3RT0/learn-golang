package main

import (
	"fmt"
)

func main() {
	var operations = []Operation{}
	operations = append(operations, Sum{1.0, 2.0})
	operations = append(operations, Sum{-1.0, 2.0})
	operations = append(operations, Subtract{1.0, 2.0})
	operations = append(operations, Divide{1.0, 2.0})
	operations = append(operations, Multiply{1.0, 2.0})

	for _, operation := range operations {
		fmt.Printf("%v = %.2f\n", operation, operation.Calculate())
	}
}

type Operation interface {
	Calculate() float64
	String() string
}

type OperationValue struct {
	value1, value2 float64
}

type Sum OperationValue

func (s Sum) Calculate() float64 {
	return s.value1 + s.value2
}

func (s Sum) String() string {
	return fmt.Sprintf("%.2f + %s", s.value1, numberToString(s.value2))
}

func numberToString(value float64) string {
	if value < 0 {
		return fmt.Sprintf("(%.2f)", value)
	}
	return fmt.Sprintf("%.2f", value)
}

type Subtract OperationValue

func (s Subtract) Calculate() float64 {
	return s.value1 - s.value2
}

func (s Subtract) String() string {
	return fmt.Sprintf("%.2f - %s", s.value1, numberToString(s.value2))
}

type Divide OperationValue

func (d Divide) Calculate() float64 {
	return d.value1 / d.value2
}

func (d Divide) String() string {
	return fmt.Sprintf("%.2f / %s", d.value1, numberToString(d.value2))
}

type Multiply OperationValue

func (m Multiply) Calculate() float64 {
	return m.value1 * m.value2
}

func (m Multiply) String() string {
	return fmt.Sprintf("%.2f x %s", m.value1, numberToString(m.value2))
}
