package main

import (
	"fmt"
	"sort"
)

func main() {
	mapStates := map[string]string{
		"SP": "São Paulo",
		"RJ": "Rio de Janeiro"}
	fmt.Println("Amount of states: ", len(mapStates))

	mapStates["SC"] = "Santa Catarina"
	fmt.Println("Amount of states: ", len(mapStates))

	var states = make(map[string]State)
	states["GO"] = State{"Goiás", 6434052, "Goiânia"}
	states["PB"] = State{"Paraíba", 3914418, "João Pessoa"}
	states["PR"] = State{"Paraná", 10997462, "Curitiba"}
	states["RN"] = State{"Rio Grande do Norte", 3373960, "Natal"}
	states["AM"] = State{"Amazonas", 3807923, "Manaus"}
	states["SE"] = State{"Sergipe", 2228489, "Aracaju"}

	fmt.Println(states)

	for _, value := range states {
		printStateInformations(value)
	}

	fmt.Println("\nDelete AM")
	delete(states, "AM")
	for _, value := range states {
		printStateInformations(value)
	}

	fmt.Println("\nPrint with alfabet order:")
	var keys []string
	for key, _ := range states {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, value := range keys {
		printStateInformations(states[value])
	}
}

func printStateInformations(state State) {
	fmt.Printf("State = %s\n", state.name)
	fmt.Printf("Capital = %s\n", state.capital)
	fmt.Printf("Population = %d\n\n", state.population)
}

type State struct {
	name       string
	population int
	capital    string
}
