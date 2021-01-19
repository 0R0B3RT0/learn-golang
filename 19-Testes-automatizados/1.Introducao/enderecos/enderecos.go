package enderecos

import "strings"

// TipoEndereco retorna o tipo do endereço
func TipoEndereco(endereco string) string {
	enderecosValidos := []string{"rua", "avenida", "estrada", "rodovia"}

	lowerAddress := strings.ToLower(endereco)
	primeiraPalavra := strings.Split(lowerAddress, " ")[0]

	for _, tipo := range enderecosValidos {
		if tipo == primeiraPalavra {
			return strings.Title(tipo)
		}
	}

	return "Não encontrado"
}
