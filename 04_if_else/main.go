package main

import "fmt"

func main() {
	idade := 18
	fmt.Println("Entre com sua idadde:")
	fmt.Scanf("%d", &idade)

	if idade >= 60 {
		fmt.Println("Olha la heim vovo, voce precisa manerar na bebedeira!")
	} else if idade >= 18 {
		fmt.Println("Legal, pode beber meu parceiro, mas nao diriga!!")
	} else {
		fmt.Println("Xiiiiii, paaa la voce precisa beber todynho")
	}
}
