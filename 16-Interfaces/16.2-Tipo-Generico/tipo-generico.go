package main

import "fmt"

func generica(valor interface{}) {
	fmt.Println(valor)
}
func main() {
	generica("String")
	generica(1)
	generica(3.4567)
	generica(true)

	mapa := map[string]interface{}{
		"nome":      "Fulado",
		"sobrenome": "de Tal",
		"idade":     38,
		"salario":   10000.32,
	}

	generica(mapa)
}
