package main

import (
	"errors"
	"fmt"
)

func main() {
	list := GenericList{1, 2, 3, 4, 5, 6}
	fmt.Println("Original list: ", list)

	removed, err := list.Remove(0)
	if err == nil {
		fmt.Println("Removed: ", removed)
		fmt.Println("List: ", list)
	} else {
		fmt.Println("Error: ", err)
	}

	removed, err = list.Remove(4)
	if err == nil {
		fmt.Println("Removed: ", removed)
		fmt.Println("List: ", list)
	} else {
		fmt.Println("Error: ", err)
	}

	removed, err = list.Remove(1)
	if err == nil {
		fmt.Println("Removed: ", removed)
		fmt.Println("List: ", list)
	} else {
		fmt.Println("Error: ", err)
	}

	removed, err = list.Remove(10)
	if err == nil {
		fmt.Println("Removed: ", removed)
		fmt.Println("List: ", list)
	} else {
		fmt.Println("Error: ", err)
	}
}

type GenericList []interface{}

func (genericList *GenericList) Remove(index int) (interface{}, error) {
	list := *genericList
	var removed interface{}
	if index < 0 || index >= len(*genericList) {
		return nil, errors.New("Invalid index")
	}

	switch index {
	case 0:
		removed = genericList.RemoveFirst()
		break
	case len(list) - 1:
		removed = genericList.RemoveLast()
		break
	default:
		removed = list[index]
		*genericList = append(list[0:index], list[index+1:len(list)]...)
		break
	}

	return removed, nil
}

func (genericList *GenericList) RemoveFirst() interface{} {
	if len(*genericList) == 0 {
		return nil
	}
	list := *genericList
	removed := list[0]
	*genericList = list[1:len(list)]

	return removed
}

func (genericList *GenericList) RemoveLast() interface{} {
	if len(*genericList) == 0 {
		return nil
	}
	list := *genericList
	removed := list[len(list)-1]
	*genericList = list[0 : len(list)-1]

	return removed
}
