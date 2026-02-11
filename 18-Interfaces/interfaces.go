package main

import (
	"fmt"
	"math"
)

// interface
type forma interface {
	area() float64
}

// "classes"/structs
type retangulo struct {
	altura  float64
	largura float64
}

type circulo struct {
	raio float64
}

// Implementa a interface
func (r retangulo) area() float64 {
	return r.altura * r.largura
}

func (c circulo) area() float64 {
	return math.Pi * math.Pow(c.raio, 2)
}

func escreverArea(f forma) {
	fmt.Printf("A área da forma é %0.2f \n", f.area())
}

func main() {
	r := retangulo{10, 15}
	escreverArea(r)

	c := circulo{20}
	escreverArea(c)
}
