package formas

import (
	"math"
	"testing"
)

func TestArea(t *testing.T) {
	t.Run("Circulo", func(t *testing.T) {
		ret := Retangulo{10, 12}
		areaEsperada := float64(120)
		areaRecebida := ret.Area()

		if areaRecebida != areaEsperada {
			t.Errorf("A área canculada não condiz com a área esperada. Esperava %f e recebeu %f",
				areaEsperada, areaRecebida)
		}
	})

	t.Run("Quadrado", func(t *testing.T) {
		circ := Circulo{(10)}

		areaEsperada := float64(math.Pi * 100)
		areaRecebida := circ.Area()

		if areaRecebida != areaEsperada {
			t.Errorf("A área canculada não condiz com a área esperada. Esperava %f e recebeu %f",
				areaEsperada, areaRecebida)
		}
	})
}
