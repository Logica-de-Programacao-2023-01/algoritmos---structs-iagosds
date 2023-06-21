package main

import "fmt"

//Crie uma struct chamada Aluno com os campos "nome", "idade" e "notas". O campo "notas" deve ser um slice de float64,
//representando as notas que o aluno tirou em uma determinada disciplina. Escreva funções que permitam adicionar ou
//remover notas do aluno, calcular a média das notas e imprimir o nome, idade e média do aluno.

type aluno struct {
	nome  string
	idade int
	notas []float64
}

func add(nt []float64) {
	var x, y int
	fmt.Print("Quais notas deseja somar? ")
	fmt.Scan(&x)
	fmt.Scan(&y)
	soma := nt[x] + nt[y]
	fmt.Println("Soma: ", soma)
}
func remove(nt []float64) []float64 {
	var i int
	var res []float64
	fmt.Print("Qual nota deseja remover? ")
	fmt.Scan(&i)
	for n := range nt {
		if nt[n] != nt[i] {
			res = append(res, nt[n])
		}
	}
	return res
}
func media(nt []float64) {
	var soma float64 = 0
	var i int = 0
	for i = range nt {
		soma += nt[i]
	}
	soma /= float64(i + 1)
	fmt.Println(soma)
}

func prt(x aluno) {
	fmt.Println("Nome: ", x.nome)
	fmt.Println("Idade: ", x.idade)
	fmt.Println("Notas: ")
	for i := range x.notas {
		fmt.Printf("\t%.2f\n", x.notas[i])
	}
}
func main() {
	iago := aluno{
		nome:  "iago",
		idade: 18,
		notas: []float64{8.5, 6.4, 7},
	}
	prt(iago)
	res := remove(iago.notas)
	fmt.Println(res)
	add(iago.notas)
	media(iago.notas)
}
