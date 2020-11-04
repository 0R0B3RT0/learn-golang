package main

import "fmt"

type ShoppingList []string

func printSlice(slice []string) {
	fmt.Println("Slice:", slice)
}
func printList(list ShoppingList) {
	fmt.Println("Shopping list:", list)
}
func main() {
	lista := ShoppingList{"Alface", "Atum", "Azeite"}
	slice := []string{"Alface", "Atum", "Azeite"}
	printSlice([]string(lista))
	printList(ShoppingList(slice))
}
