package main

import "fmt"

func main() {
	// Null value in go = 0
	var x int
	fmt.Println("x = ", x)

	var y = 10
	fmt.Println("y = ", y)

	var z, w, v float64

	fmt.Printf("z = %.2f | w = %.2f | v = %.2f \n", z, w, v)

	var name, lastname string = "Joao Victor", "Rocha"
	fmt.Println(name, lastname)
}
