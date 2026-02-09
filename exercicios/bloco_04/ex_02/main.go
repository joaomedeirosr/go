package main

import "fmt"

func main() {

	var soma float64 // Go lang variaveis sao 0 por default, chamados de valores neutros ou Zero Values
	for i := 1; i <= 4; i++ {
		var altura float64
		fmt.Println("Digite a sua altura")
		fmt.Scanf("%f", &altura)
		soma += altura
	}

	fmt.Printf("A soma de todas as alturas e: %.2f m \n", soma)
}
