package main

import "fmt"

func main() {

	var option int

	fmt.Printf("Seleciona um opcao: ")
	fmt.Println(`1 - Banana, 2 - Maca, 3 - Maracuja, 4 - Pera`)
	fmt.Scanf("%d", &option)

	switch option {
	case 1:
		fmt.Println("Banana")
	case 2:
		fmt.Println("Maca")
	case 3:
		fmt.Println("Maracuja")
	case 4:
		fmt.Println("Pera")
	}
}
