package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQuicksort(t *testing.T) {
	var tests = []Test{
		{[]int{3, 2, 1}, []int{1, 2, 3}},
		{[]int{300, 22, 150}, []int{22, 150, 300}},
	}

	for _, test := range tests {
		actual := quicksort(test.input)
		assert.Equal(t, actual, test.output, "they should be equal")
	}

}

type Test struct {
	input  []int
	output []int
}
