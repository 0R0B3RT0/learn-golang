package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

func (u usuario) salvar() {
	fmt.Printf("Salvando o usuário %s no banco de dados\n", u.nome)
}

func (u usuario) maiorDeIdade() bool {
	return u.idade >= 18
}

func (u *usuario) fazerAniversario() {
	u.idade++
}

func main() {
	fmt.Println("Métodos")

	usuario1 := usuario{"Fulano", 23}
	usuario1.salvar()
	fmt.Printf("É maior de idade? %s\n", usuario1.maiorDeIdade())
	fmt.Printf("Idade: %d\n", usuario1.idade)
	usuario1.fazerAniversario()
	fmt.Printf("Idade após aniversário: %d\n", usuario1.idade)

	usuario{"Ciclano", 31}.salvar()
}
