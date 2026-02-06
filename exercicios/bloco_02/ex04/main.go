package main

import (
	"fmt"
)

func main() {
	var number float32
	fmt.Println("Entre com numero")
	fmt.Scanf("%f", &number)
	double := number * 2

	fmt.Printf("O dobro do numero %.2f e: %.2f \n", number, double)
}
