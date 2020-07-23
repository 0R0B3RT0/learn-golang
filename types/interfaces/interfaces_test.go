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
	operations = append(operations, OperationTest{Sum{-1, 2}, 1, "(-1.00) + 2.00"})
	operations = append(operations, OperationTest{Sum{-1, -5}, -6, "(-1.00) + (-5.00)"})

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
