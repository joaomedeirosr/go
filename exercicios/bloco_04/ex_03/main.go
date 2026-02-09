package main

import (
	"fmt"
	"strconv"
)

func main() {

	var soma float64

	for {
		var input string
		fmt.Println("Digite um valor: ")
		fmt.Scanf("%s", &input)

		if input == "" {
			break
		}
		value, err := strconv.ParseFloat(input, 64)

		if err != nil {
			fmt.Println("Entre com um valor valido")
			continue
		}

		soma += value
	}
	fmt.Printf("O valor total e R$: %.2f \n", soma)
}
