package main

import "fmt"

func main() {
	i := 1
	for i <= 10 {
		fmt.Println("Outra maneira de for", i)
		i++
	}
}

// Podemos trabalhar no Go, com for i true .. e condicoes enfim tudo que
// as outras linguagens tambem tem, aqui o Go faz o for se transformar em
// um while porque Go, nao possui while

// Fazer esse codigo acima e algo como while <= 10 entao ..
