package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSoma(t *testing.T) {
	var operations = []OperationTest{}
	operations = append(operations, OperationTest{Sum{1.0, 1.0}, 2, "1.00 + 1.00"})
	operations = append(operations, OperationTest{Sum{0.5, 0.5}, 1, "0.50 + 0.50"})
	operations = append(operations, OperationTest{Sum{1, 1}, 2, "1.00 + 1.00"})
	operations = append(operations, OperationTest{Sum{10, -5}, 5, "10.00 + (-5.00)"})
	operations = append(operations, OperationTest{Sum{-1, 2}, 1, "-1.00 + 2.00"})
	operations = append(operations, OperationTest{Sum{-1, -5}, -6, "-1.00 + (-5.00)"})

	for _, operation := range operations {
		assert.Equal(t, operation.expected, operation.Calculate(), "they should be equal")
		assert.Equal(t, operation.operationString, operation.Operation.String(), "they should be equal")
	}
}

func TestSubtract(t *testing.T) {
	var operations = []OperationTest{}
	operations = append(operations, OperationTest{Subtract{1.0, 1.0}, 0, "1.00 - 1.00"})
	operations = append(operations, OperationTest{Subtract{0.5, 0.5}, 0, "0.50 - 0.50"})
	operations = append(operations, OperationTest{Subtract{1, 2}, -1, "1.00 - 2.00"})
	operations = append(operations, OperationTest{Subtract{10, -5}, 15, "10.00 - (-5.00)"})
	operations = append(operations, OperationTest{Subtract{-1, 2}, -3, "-1.00 - 2.00"})
	operations = append(operations, OperationTest{Subtract{-1, -5}, 4, "-1.00 - (-5.00)"})

	for _, operation := range operations {
		assert.Equal(t, operation.expected, operation.Calculate(), "they should be equal")
		assert.Equal(t, operation.operationString, operation.Operation.String(), "they should be equal")
	}
}

func TestMultiply(t *testing.T) {
	var operations = []OperationTest{}
	operations = append(operations, OperationTest{Multiply{1.0, 1.0}, 1, "1.00 x 1.00"})
	operations = append(operations, OperationTest{Multiply{0.5, 0.5}, 0.25, "0.50 x 0.50"})
	operations = append(operations, OperationTest{Multiply{1, 0.5}, 0.5, "1.00 x 0.50"})
	operations = append(operations, OperationTest{Multiply{10, -5}, -50, "10.00 x (-5.00)"})
	operations = append(operations, OperationTest{Multiply{-1, 2}, -2, "-1.00 x 2.00"})
	operations = append(operations, OperationTest{Multiply{-1, -5}, 5, "-1.00 x (-5.00)"})

	for _, operation := range operations {
		assert.Equal(t, operation.expected, operation.Calculate(), "they should be equal")
		assert.Equal(t, operation.operationString, operation.Operation.String(), "they should be equal")
	}
}

func TestDivide(t *testing.T) {
	var operations = []OperationTest{}
	operations = append(operations, OperationTest{Divide{1.0, 1.0}, 1, "1.00 / 1.00"})
	operations = append(operations, OperationTest{Divide{0.5, 0.5}, 1, "0.50 / 0.50"})
	operations = append(operations, OperationTest{Divide{1, 0.5}, 2, "1.00 / 0.50"})
	operations = append(operations, OperationTest{Divide{10, -5}, -2, "10.00 / (-5.00)"})
	operations = append(operations, OperationTest{Divide{-1, 2}, -0.5, "-1.00 / 2.00"})
	operations = append(operations, OperationTest{Divide{-1, -5}, 0.20, "-1.00 / (-5.00)"})

	for _, operation := range operations {
		assert.Equal(t, operation.expected, operation.Calculate(), "they should be equal")
		assert.Equal(t, operation.operationString, operation.Operation.String(), "they should be equal")
	}
}

type OperationTest struct {
	Operation
	expected        float64
	operationString string
}
