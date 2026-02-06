package main

import "fmt"

func main() {
	fmt.Println("Ola, bom dia qual e o seu nome?")
	var nome string
	fmt.Scanf("%s", &nome)
	fmt.Println("E um prazer te conhecer: ", nome)
}
