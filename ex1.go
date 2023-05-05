package main

import (
	"fmt"
	"math"
)

type circulo struct {
	raio float64
}

func area(c circulo) float64 {
	a := (c.raio * c.raio) * math.Pi
	return a
}
func main() {
	r := circulo{raio: 2}
	fmt.Printf("A área do círculo é: %.2f", area(r))
}
