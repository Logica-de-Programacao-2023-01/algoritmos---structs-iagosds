package main

import "fmt"

type info struct {
	t   string
	a   string
	ano string
}

func livro(l info) {
	fmt.Println("t:", l.t)
	fmt.Println("a: ", l.a)
	fmt.Println("ano: ", l.ano)
}
func main() {
	p := info{t: "titulo", a: "autor", ano: "2023"}
	livro(p)

}
