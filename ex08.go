package main

import "fmt"

//Crie uma struct chamada Viagem com os campos "origem", "destino", "data" e "preço". Escreva uma função que receba um
//slice de Viagens como parâmetro e retorne a viagem mais cara.

type viagem struct {
	origem  string
	destino string
	data    int
	preco   int
}

func maiscara(viajar []viagem) viagem {
	maior := 0
	for i := range viajar {
		if viajar[i].preco > maior {
			maior = viajar[i].preco
		}
	}
	for i := range viajar {
		if maior == viajar[i].preco {
			return viajar[i]
		}
	}
	return viajar[1]
}
func main() {
	vias := []viagem{
		{
			origem:  "x",
			destino: "y",
			data:    15,
			preco:   1000,
		},
		{
			origem:  "y",
			destino: "y",
			data:    16,
			preco:   1100,
		},
	}
	res := maiscara(vias)
	fmt.Println(res)
}
