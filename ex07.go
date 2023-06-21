package main

import "fmt"

//Crie uma struct chamada Animal com os campos "nome", "espécie", "idade" e "som". Escreva funções que permitam
//modificar o som que o animal faz e uma função que imprima as informações do animal e o som que ele faz.

type animal struct {
	nome    string
	especie string
	idade   int
	som     string
}

func trocasom(s string) {
	res := "N"
	fmt.Println("Deseja alterar o som do animal?(s/n)")
	fmt.Scan(&res)
	if res == "s" || res == "S" {
		fmt.Print("Qual som deseja que o animal faça? ")
		fmt.Scanln(&s)
	}
}
func prt(x animal) {
	fmt.Println(x)
}
func main() {
	leao := animal{
		nome:    "leo",
		especie: "leao",
		idade:   10,
		som:     "raur!",
	}
	prt(leao)
	trocasom(leao.som)
	prt(leao)

}
