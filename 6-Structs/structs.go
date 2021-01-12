package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint8
}

func main() {
	fmt.Println("Arquivo structs")

	var u usuario
	u.nome = "Fulano"
	u.idade = 18
	fmt.Println(u)

	e := endereco{"Rua Um", 201}
	u2 := usuario{"João", 30, e}
	fmt.Println(u2)

	u3 := usuario{nome: "Luiz"}
	fmt.Println(u3)

	u4 := usuario{idade: 30}
	fmt.Println(u4)

	u5 := usuario{"José", 30, endereco{"Rua Um", 201}}
	fmt.Println(u5)
}
