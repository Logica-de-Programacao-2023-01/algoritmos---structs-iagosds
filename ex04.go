package main

import "fmt"

type musicas struct {
	titulo  string
	artista string
	duracao float64
}
type playlist struct {
	nome string
	msk  []musicas
}

func info(p playlist) {
	fmt.Printf("\tPLAYLIST\n")
	fmt.Println("---------------------")
	fmt.Println("Nome: ", p.nome)
	var tempo float64 = 0
	for i := range p.msk {
		tempo += p.msk[i].duracao
	}
	fmt.Printf("Duração: %.2f min\n\n", tempo)
	for i := range p.msk {
		fmt.Printf("\tMÚSICA %d\n", i+1)
		fmt.Println("---------------------")
		fmt.Println("Título: ", p.msk[i].titulo)
		fmt.Println("Artista: ", p.msk[i].artista)
		fmt.Printf("Duração: %.2f min\n\n", p.msk[i].duracao)
	}
}
func main() {
	slic := []musicas{
		{
			titulo:  "ttl",
			artista: "x",
			duracao: 3.30,
		},
		{
			titulo:  "ttl2",
			artista: "y",
			duracao: 4.30,
		},
	}
	play := playlist{
		nome: "Playlist do SDS",
		msk:  slic,
	}
	info(play)
}
