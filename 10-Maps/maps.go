package main

import "fmt"

func main() {
	fmt.Println("Maps")

	usuario := map[string]string{
		"nome":      "João",
		"sobrenome": "da Silva",
	}

	fmt.Println(usuario)
	fmt.Println(usuario["nome"])
	fmt.Println(usuario["sobrenome"])
	fmt.Println(usuario["algumacoisa"])

	delete(usuario, "nome")
	fmt.Println(usuario)

}
