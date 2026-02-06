package main

import (
	"fmt"
	"math"
)

func main() {
	var number int
	fmt.Print("Entre com um numero:")
	fmt.Scanf("%d", &number)

	result := math.Sqrt(float64(number))
	fmt.Printf("A raiz quadrade de %d e: %.2f \n", number, result)
}
