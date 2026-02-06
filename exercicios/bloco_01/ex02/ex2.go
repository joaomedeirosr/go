package main

import "fmt"

func main() {
	fmt.Println("Table with (AND) Logic:")
	fmt.Println("true && false = ", true && false)
	fmt.Println("true && true = ", true && true)
	fmt.Println("false && true = ", false && true)
	fmt.Println("false && false =", false && false)

	fmt.Println("#---------------------------#")

	fmt.Println("Table with (OR) Logic:")
	fmt.Println("false || false = ", false || false)
	fmt.Println("false || true = ", false || true)
	fmt.Println("true || false =", true || false)
	fmt.Println("true || true = ", true || true)

}
