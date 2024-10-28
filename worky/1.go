package main

import (
	"fmt"
	"math"
)

func main() {
	var lado1, lado2 float64
	const PRECISION = 2
	fmt.Print("Ingrese lado 1: ")
	fmt.Scanln(&lado1)
	fmt.Print("Ingrese lado 2: ")
	fmt.Scanln(&lado2)
	area := (lado1 * lado2) / 2
	hipotenusa := math.Sqrt(math.Pow(lado1,2) + math.Pow(lado2, 2))
	perimetro := lado1 + lado2 + hipotenusa

	//fmt.Printf("\nArea: %.2f", area)
	//fmt.Printf("\nPerimetro: %.2f", perimetro)
	fmt.Printf("\nArea: %.*f", PRECISION, area)
	fmt.Printf("\nPerimetro: %.*f", PRECISION, perimetro)
}