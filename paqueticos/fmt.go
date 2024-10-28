package main

import "fmt"

func main() {
	fmt.Println("Wembie") //Hace salto de linea
	fmt.Print("Wombo") //No hace salto de linea
	gamerName := "Wembie"
	fmt.Printf("\nHola, me llaman %s y soy demasiado pro UwU.\n", gamerName)
	greeting := fmt.Sprintf("\nHola, me llaman %s y soy demasiado pro UwU | Soy greeting", gamerName)
	fmt.Println(greeting)

	var(
		name string
		age int
	)
	fmt.Print("Ingrese tu nombre pa: ")
	fmt.Scanln(&name) //Ejemplo: Si pone JUAN, melo, pero si pone JUAN ESTEBAN, no daria tocaria con otra variable como name2 quedaria: fmt.Scanln(&name, &name2)
	fmt.Print("Ingrese tu edad pa: ")
	fmt.Scanln(&age)
	saludo := fmt.Sprintf("Que se dice %s por ahi vi que tiene %d añitos", name, age) //si no sabes que tipo de dato va a hacer especificamente colocar %v
	fmt.Println(saludo)
	fmt.Printf("El tipo de name es: %T\n", name)
}