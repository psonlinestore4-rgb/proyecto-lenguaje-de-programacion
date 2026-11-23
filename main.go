package main

import (
	"fmt"
	"math"
)

type Figura interface {
	Area() float64
	Perimetro() float64
	Nombre() string
}

type Rectangulo struct {
	Largo, Ancho float64
}

func (r Rectangulo) Area() float64      { return r.Largo * r.Ancho }
func (r Rectangulo) Perimetro() float64 { return 2 * (r.Largo + r.Ancho) }
func (r Rectangulo) Nombre() string     { return "Rectángulo" }

type Circulo struct {
	Radio float64
}

func (c Circulo) Area() float64      { return math.Pi * c.Radio * c.Radio }
func (c Circulo) Perimetro() float64 { return 2 * math.Pi * c.Radio }
func (c Circulo) Nombre() string     { return "Círculo" }

type Triangulo struct {
	A, B, C float64
}

func (t Triangulo) Area() float64 {
	s := (t.A + t.B + t.C) / 2
	return math.Sqrt(s * (s - t.A) * (s - t.B) * (s - t.C))
}

func (t Triangulo) Perimetro() float64 { return t.A + t.B + t.C }
func (t Triangulo) Nombre() string     { return "Triángulo" }

func ImprimirFigura(f Figura) {
	fmt.Printf("Figura: %s\n", f.Nombre())
	fmt.Printf(" Área: %.2f\n", f.Area())
	fmt.Printf(" Perímetro: %.2f\n\n", f.Perimetro())
}

func main() {
	figuras := []Figura{
		Rectangulo{Largo: 3, Ancho: 4},
		Circulo{Radio: 5},
		Triangulo{A: 3, B: 4, C: 5},
	}

	for _, f := range figuras {
		ImprimirFigura(f)
	}
}