package main

import (
	"fmt"
	"math/rand"
)

func main() {
	//rand.Seed(time.Now().UnixNano()) //Por si no se generar varios aleatorios
	//fmt.Println(rand.Intn(100))
	play()

}

func play(){
	random_int := rand.Intn(100)
	var num_ingresado, intentos int
	const MAX_TRY = 10
	for intentos < MAX_TRY {
		fmt.Printf("Ingresa un numero (intentos restantes: %d): ", MAX_TRY - intentos)
		fmt.Scanln(&num_ingresado)
		if num_ingresado == random_int {
			fmt.Println("Gano mij@")
			try_again()
			return
		} else if num_ingresado < random_int {
			fmt.Println("Grave pa, el numero ese es mayor mij@")
		} else if num_ingresado > random_int {
			fmt.Println("Grave pa, el numero ese es menor mij@")
		}
		intentos++
	}
	fmt.Println("Mucha pichurria, muy pailaaaaaa uy no... muy suaveee como pierde xd, mero panfleto, el numero que estaba buscando era el:", random_int)
	try_again()

}

func try_again(){
	var eleccion string
	fmt.Print("Quiere jugar de nuevo o no hable claro (s/n): ")
	fmt.Scanln(&eleccion)
	switch eleccion {
	case "s":
		play()
	case "n":
		fmt.Println("Abrite entonces")
	default:
		fmt.Println("Como digita mal eso, severa gueva, try again ome buñuelo")
		try_again()
	}
}