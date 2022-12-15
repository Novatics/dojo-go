package main

import "fmt"

func main() {

	fmt.Print("Digite o valor \n")
	var saqueValor int

	// Taking input from user
	fmt.Scanln(&saqueValor)

	calculaSaque(&saqueValor)
	fmt.Print(saqueValor)

}

func calculaSaque(int &valor) {
	valor = valor / 2
}
