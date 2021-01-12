package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    uint8
}

type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {
	fmt.Println("Herança")

	p1 := pessoa{"João", "da Silta", 20, 170}
	fmt.Println(p1)

	e1 := estudante{p1, "Engenharia", "USP"}
	fmt.Println(e1)

	fmt.Println("Pessoa: ", e1.pessoa)

	fmt.Println("nome: ", e1.nome)
	fmt.Println("sobrenome: ", e1.sobrenome)
	fmt.Println("idade: ", e1.idade)
	fmt.Println("altura: ", e1.altura)
	fmt.Println("curso: ", e1.curso)
	fmt.Println("faculdade: ", e1.faculdade)
}
