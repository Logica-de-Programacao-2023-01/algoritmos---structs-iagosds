package main

import "fmt"

//Crie uma struct chamada Filme com os campos "título", "diretor", "ano" e "avaliações". O campo "avaliações" deve ser
//um slice de inteiros representando as notas que o filme recebeu dos espectadores. Escreva funções que permitam
//adicionar ou remover avaliações do filme, calcular a média das avaliações e imprimir as informações do filme e sua
//média de avaliações.

type filme struct {
	titulo     string
	diretor    string
	ano        int
	avaliacoes []int
}

func media(av []int) {
	var soma float64 = 0
	for i := range av {
		soma += float64(av[i])
	}
	fmt.Println(soma)
	soma /= float64(len(av))
	fmt.Printf("%.2f\n", soma)

}

func add(av []int) {
	var nota int
	fmt.Print("Qual nota deseja atribuir ao filme? ")
	fmt.Scan(&nota)
	av = append(av, nota)
	fmt.Println(av)
}
func remove(av []int) {
	var nota int
	fmt.Print("Qual nota deseja remover? ")
	fmt.Scan(&nota)
	if nota == len(av) {
		av = append(av[:nota])
	} else if nota == 0 {
		av = append(av[(nota + 1):])
	} else {
		av = append(av[:nota], av[nota+1:]...)
	}
	fmt.Println(av)

}
func prt(movie filme) {
	fmt.Println("Titulo: ", movie.titulo)
	fmt.Println("Diretor: ", movie.diretor)
	fmt.Println("Ano: ", movie.ano)
	fmt.Println("Avalicações: ")
	for _, i := range movie.avaliacoes {
		fmt.Printf("\t%d\n", i)
	}
}
func main() {
	flm := filme{
		titulo:     "Super Mário",
		diretor:    "Tarantino",
		ano:        2023,
		avaliacoes: []int{8, 3, 6, 9},
	}
	add(flm.avaliacoes)
	remove(flm.avaliacoes)
	media(flm.avaliacoes)
	prt(flm)
}
