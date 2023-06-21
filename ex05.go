package main

type musicas struct {
	titulo  string
	artista string
	duracao float64
}
type playlist struct {
	nome string
	msk  []musicas
}

func search(mus string)[]playlist{
	for i:= range
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
}
