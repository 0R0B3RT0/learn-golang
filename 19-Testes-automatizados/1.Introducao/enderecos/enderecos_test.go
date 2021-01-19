package enderecos

import (
	"testing"
)

type cenarioDeTeste struct {
	endereco string
	esperado string
}

func TestTipoEndereco(t *testing.T) {
	cenariosDeTeste := []cenarioDeTeste{
		{"Avenida Brasil", "Avenida"},
		{"Rua Brasil", "Rua"},
		{"Rodovia Brasil", "Rodovia"},
		{"Estrada Brasil", "Estrada"},
		{"ESTRADA Brasil", "Estrada"},
		{"rua", "Rua"},
		// {"Praça Brasil", "Não encontrado"},
		// {"", "Não encontrado"},
	}

	for _, cenario := range cenariosDeTeste {
		saida := TipoEndereco(cenario.endereco)

		if saida != cenario.esperado {
			t.Errorf("O endereço recebido é diferente do esperado. Esperava %s e recebeu %s!",
				cenario.esperado, saida)
		}
	}

}
