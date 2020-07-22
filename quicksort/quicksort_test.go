package main

import (
	"reflect"
	"testing"
)

func TestQuicksort(t *testing.T) {
	var tests = []Test{
		{[]int{3, 2, 1}, []int{1, 2, 3}},
		{[]int{300, 22, 150}, []int{22, 150, 300}},
	}

	for _, test := range tests {
		actual := quicksort(test.input)
		if !reflect.DeepEqual(actual, test.output) {
			t.Error(
				"For input: ", test.input,
				"expected:", test.output,
				"got:", actual)
		}
	}

}

type Test struct {
	input  []int
	output []int
}
