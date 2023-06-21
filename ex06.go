package main

import "fmt"

type funcionario struct {
	nome    string
	salario float64
	idade   int
}

func sal(s float64) {
	var pctg float64
	fmt.Print("Em quantos % será aumentado o salário do funcionário? ")
	fmt.Scan(&pctg)
	pctg = (pctg / 100) + 1
	s = s * pctg
	fmt.Printf("Novo salário: %.2f", s)
}
func tempo(id int) (int, error) {
	if id >= 18 {
		id -= 18
	} else if id < 18 {
		return 0, fmt.Errorf("Funcionário não tem idade para trabalhar")
	}
	return id, nil
}
func main() {
	iago := funcionario{
		nome:    "iago",
		salario: 1200,
		idade:   15,
	}
	res, err := tempo(iago.idade)
	if err != nil {
		fmt.Println(err)
	} else if res == 0 {
		fmt.Println("Funcionário começou este ano.")
	} else {
		fmt.Println("Funcionário trabalha a ", res, " anos")
		sal(iago.salario)
	}

}
