package main

import (
	"fmt"
)

func main() {
	hello()
	hello_name("Wembie")
	saluditoo := hello_return("Juancho")
	fmt.Println(saluditoo)
	fmt.Println(suma(60,9))
	sumita, multiplicacioncita := calculadora(4, 5)
	fmt.Printf("La sumita es: %d, y la multiplicacioncita es: %d", sumita, multiplicacioncita)
}

func hello(){
	fmt.Println("Que se dice el/la soci@")
}

func hello_name( name string ){
	fmt.Printf("Que se dice %s", name)
}

func hello_return( name string ) string{
	return "\nQue se dice " + name + ", bien o q"
}

func suma( a, b int ) int {
	return a + b
}

/*func calculadora( a, b int )( int, int ){
	sum := a + b
	mult := a * b
	return sum, mult
}*/
func calculadora( a, b int )( sum, mult int ){
	sum = a + b
	mult = a * b
	return
}