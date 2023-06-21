package main

import "fmt"

type triangulo struct {
	base   float64
	altura float64
}

func area(x triangulo) {
	ar := (x.base * x.altura) / 2
	fmt.Print(ar)
}
func main() {
	tri := triangulo{
		base:   10,
		altura: 15,
	}
	area(tri)
}
