package main

import "fmt"

func main() {
	fmt.Println("Posso digitar o que eu quiser em Stirngs de aspas duplas")

	fmt.Println(`Aqui eu posso digitar strings
	Multiline sem problemas!,
	Tudo numa boa`)

	fmt.Println("O nome Joao tem:", len("Joao"), "letras")

	// Acessando sub-strings
	fmt.Println(string("Joao"[0]))

	// Byte e representado por um inteiro de 1 byte = 8 bits (int)
}
