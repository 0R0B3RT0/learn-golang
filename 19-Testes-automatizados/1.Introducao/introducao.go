package main

import (
	"fmt"

	"introducao-testes/enderecos"
)

func main() {

	fmt.Println(enderecos.TipoEndereco("rUa do joão"))
	fmt.Println(enderecos.TipoEndereco("RODOVIA  ESTADUAL"))
	fmt.Println(enderecos.TipoEndereco("AVENIDA XV DE NOVEMBRO"))
	fmt.Println(enderecos.TipoEndereco("estrada da capela"))
	fmt.Println(enderecos.TipoEndereco("viela da capela"))
}
