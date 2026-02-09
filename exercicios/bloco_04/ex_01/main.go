package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Digite uma palavra")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	word := scanner.Text()
	fmt.Println("A palavra digitada foi:", word)

	var counter int
	for i := range word {
		if word[i] == 'a' || word[i] == 'A' {
			counter += 1
		}
	}

	fmt.Printf("A palavra %s: tem %d letras 'A' \n", word, counter)

}
