package formas

import (
	"fmt"
	"math"
)

type Forma interface {
	area() float64
}

type Retangulo struct {
	altura  float64
	largura float64
}

type Circulo struct {
	raio float64
}

func (r Retangulo) Area() float64 {
	return r.altura * r.largura
}

func (c Circulo) Area() float64 {
	return math.Pi * math.Pow(c.raio, 2)
}

func CalculaArea(f Forma) {
	fmt.Printf("A área da forma é %0.2f\n", f.area())
}
