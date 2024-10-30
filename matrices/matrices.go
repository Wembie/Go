package main

import "fmt"

func main() {
	var matriz [5]int
	matriz[0] = 10
	matriz[1] = 10
	fmt.Println(matriz)
	a := [5]int{1,2,3,4,5} //fijos
	fmt.Println(a)
	b := [...]int{1,3,5,7,9} //no se sabe cuantos elementos va a tener
	fmt.Println(b)
	for i:=0; i < len(b); i++{
		fmt.Println(b[i])
	}
	for index, value := range b{ //index y valor
		fmt.Printf("Indice: %d, Valor: %d\n", index, value)
	}
	for _, value := range b{ //solo valor
		fmt.Printf("Indice: Paila, Valor: %d\n", value)
	}
	//filas,columnas
	var matrix = [3][3]int{{1,2,3},{4,5,6},{7,8,9}}
	fmt.Println(matrix)
}