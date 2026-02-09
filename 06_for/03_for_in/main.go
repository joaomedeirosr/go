package main

import "fmt"

func main() {

	nome := "Joao Victor Rocha"

	// For detalhado

	for i := range nome {
		fmt.Println(i, nome[i], string(nome[i]))
	}

	for i := range nome {
		fmt.Println(string(nome[i]))
	}

	for i, v := range nome {
		fmt.Println(i, string(v))
	}

	/*
			For com "_ (underline)" ele serve para falar que nao vamos usar vai servir so pra mostrar o indice da
		  	estrutura de dados
	*/

	for _, v := range nome {
		fmt.Println(string(v))
	}
}
