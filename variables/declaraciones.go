package main

import ( 
	"fmt"
	"rsc.io/quote"
)

//CONSTANTES
//const pi float32 = 3.1416
const PI = 3.1416

const (
	X = 100
	Y = 0b1010 //binario
	Z = 0o12 //octal
	W = 0xFF //Hexadecimal
)

/*const (
	DOMINGO = 1
	LUNES = 2
	MARTES = 3
	MIERCOLES = 4
	JUEVES = 5
	VIERNES = 6
	SABADO = 7
)*/

// iota iniciar desde 0 y los demas valores incrementando de a 1
const (
	DOMINGO = iota + 1
	LUNES
	MARTES
	MIERCOLES
	JUEVES
	VIERNES
	SABADO 
)
func main() {
	fmt.Println("Hola Wumpux")
	fmt.Println(quote.Go())

	//Variables
	/*var name string
	var first_name, second_name string
	var age int

	name = "Wembie"
	first_name = "Juan"
	second_name = "Acosta"
	age = 22

	var (
		name, first_name, second_name string
		age int
	)

	var (
		name string = "Wembie"
		first_name string = "Juan"
		second_name string = "Acosta"
		age int = 22
	)	
	
	//No es necesario asignarle el tipo de dato
	var (
		name = "Wembie"
		first_name = "Juan"
		second_name  = "Acosta"
		age = 22
	)
	
	var name, first_name, second_name, age = "Wembie", "Juan", "Acosta", "22"
	*/
	//:= sirve para declarar una variable nueva y inicializarla, solamente dentro de las funciones y con var si se puede fuera de las funciones, para modificar si seria normalmente con =
	name, first_name, second_name, age := "Wembie", "Juan", "Acosta", "22"
	fmt.Println(name, first_name, second_name, age)
	fmt.Println(PI)
	fmt.Println(X,Y,Z,W)
	fmt.Println(DOMINGO, SABADO)
}