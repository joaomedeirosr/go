package main

import "fmt"

/*
	Tipos no go:
	- Inteiros: int8,int16,int32,int64
	- Inteiros positivos: uint8,uint16,uint32,uint64

	Float:
	float32, float64
*/

func main() {
	fmt.Println("1+1 =", 1+1)
	fmt.Println("1+ 2.0 =", 1+2.5)
	fmt.Println("10 / 3 =", 10/3)
	fmt.Println("10 / 3 =", 10/3.0)
	fmt.Println("100*0.75 =", 100*0.75)

	fmt.Printf("Tipo de 100 * 0.75 = %T \n", 100*0.75)

	fmt.Println("5 % 3 =", 5%3)

}
