package main

import "fmt"

func main() {

	soma := 0.0
	for i := 1; i <= 4; i++ {
		var altura float64
		fmt.Println("Digite a sua altura")
		fmt.Scanf("%f", &altura)
		soma += altura
	}

	fmt.Printf("A soma de todas as alturas e: %.2f \n", soma)
}
