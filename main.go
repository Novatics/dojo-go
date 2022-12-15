package main

import "fmt"

type register struct {
	nota100 int
	nota50 int
	nota20 int
	nota10 int
}

func main() {

	fmt.Print("Digite o valor \n")
	var saqueValor int

	// Taking input from user
	fmt.Scanln(&saqueValor)
	// s := register{nota100: 100, nota50: 100, nota20: 100, nota10: 100}

	calculaSaque(100)
	fmt.Print(saqueValor)

}

func calculaSaque(valor int) int {
	// nota100 := valor / 100


	// nota50 := valor / 50

	// nota20 := valor / 20

	// nota10 := valor /10
	return valor
}
