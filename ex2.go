package main

import "fmt"

type end struct {
	rua    string
	numero int
	cidade string
	estado string
}

type pessoa struct {
	nome     string
	idade    int
	endereco end
}

func main() {
	p := pessoa{
		nome:  "iago",
		idade: 18,
		endereco: end{
			rua:    "x",
			numero: 001,
			cidade: "brasilia",
			estado: "df",
		},
	}
	fmt.Println("nome: ", p.nome)
	fmt.Println("idade: ", p.idade)
	fmt.Println("Endereço:")
	fmt.Printf("\trua: %s\n", p.endereco.rua)
	fmt.Printf("\tnumero: %d\n", p.endereco.numero)
	fmt.Printf("\tcidade: %s\n", p.endereco.cidade)
	fmt.Printf("\testado: %s", p.endereco.estado)
}
