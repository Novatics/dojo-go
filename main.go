package main

import "fmt"

type register struct {
	nota100 int
	nota50  int
	nota20  int
	nota10  int
}

func main() {

	fmt.Print("Digite o valor \n")
	var saqueValor int

	// Taking input from user
	fmt.Scanln(&saqueValor)

	calculaSaque(220)
	fmt.Print(saqueValor)

}

func calculaResto(nota int, valor int) int {

    result := valor % nota

    return result
}

func numeroDeNotas(nota int, valor int) int {
    return valor / nota
}

func calculaSaque(valor int) register {
	s := register{nota100: 0, nota50: 0, nota20: 0, nota10: 0}
	// resto := 0

	s.nota100 = numeroDeNotas(100, valor) 
	valor = calculaResto(100,valor)
	s.nota50 = numeroDeNotas(50, valor) 
	valor = calculaResto(50,valor)
	s.nota20 = numeroDeNotas(20, valor) 
	valor = calculaResto(20,valor)
	s.nota10 = numeroDeNotas(10, valor) 
	valor = calculaResto(10,valor)

	fmt.Print("AAAAAAAAAAA")
	fmt.Print(s.nota100)

	return s
}
