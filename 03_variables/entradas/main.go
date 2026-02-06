package main

import "fmt"

func main() {
	fmt.Println("Enter with one value, please:")

	var x float64
	fmt.Scanf("%f", &x)

	fmt.Println("Voce escolheu o valor: ", x)
}
