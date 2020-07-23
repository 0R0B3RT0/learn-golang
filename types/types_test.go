package main

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRemoveFirst(t *testing.T) {
	testList := []GenericListTest{
		{GenericList{1, 2, 3}, GenericList{2, 3}, 1},
		{GenericList{2, 3}, GenericList{3}, 2},
		{GenericList{3}, GenericList{}, 3},
		{GenericList{}, GenericList{}, nil},
	}

	for _, test := range testList {
		removed := test.input.RemoveFirst()
		assert.Equal(t, test.removed, removed, "they should be equal")
		assert.Equal(t, test.expected, test.input, "they should be equal")
	}
}

func TestRemoveLast(t *testing.T) {
	var testList = []GenericListTest{
		{GenericList{1, 2, 3}, GenericList{1, 2}, 3},
		{GenericList{1, 2}, GenericList{1}, 2},
		{GenericList{1}, GenericList{}, 1},
		{GenericList{}, GenericList{}, nil},
	}

	for _, test := range testList {
		removed := test.input.RemoveLast()
		assert.Equal(t, test.removed, removed, "they should be equal")
		assert.Equal(t, test.expected, test.input, "they should be equal")
	}
}

func TestRemove(t *testing.T) {
	var testList = []GenericListTestRemove{
		{GenericList{1, 2, 3}, GenericList{1, 3}, 2, 1},
		{GenericList{1, 3}, GenericList{1}, 3, 1},
		{GenericList{1}, GenericList{}, 1, 0},
	}

	for _, test := range testList {
		removed, _ := test.input.Remove(test.index)
		assert.Equal(t, test.removed, removed, "they should be equal")
		assert.Equal(t, test.expected, test.input, "they should be equal")
	}
}

func TestMustBeErrorWhenRemoveInvalidIndex(t *testing.T) {
	list := GenericList{1, 2, 3}

	removed, err := list.Remove(10)

	assert.Nil(t, removed, "they should be null")
	assert.NotNil(t, err, "they should not be null")
	assert.Equal(t, errors.New("Invalid index"), err, "they should be equal")
}

func TestMustBeErrorWhenRemoveNegativeIndex(t *testing.T) {
	list := GenericList{1, 2, 3}

	removed, err := list.Remove(-1)

	assert.Nil(t, removed, "they should be null")
	assert.NotNil(t, err, "they should not be null")
	assert.Equal(t, errors.New("Invalid index"), err, "they should be equal")

}

type GenericListTest struct {
	input    GenericList
	expected GenericList
	removed  interface{}
}

type GenericListTestRemove struct {
	input    GenericList
	expected GenericList
	removed  interface{}
	index    int
}
