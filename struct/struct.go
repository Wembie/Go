package main

import "fmt"

type Persona struct {
	nombre string
	edad int
	correo string
}

func main() {
	/*var personita Persona

	personita.nombre = "Wembie"
	personita.edad = 22
	personita.correo = "Awitadecoco@hotmail.com"*/
	personita := Persona{"Wembie", 22, "Awitadecoco@hotmail.com"}
	fmt.Println(personita)
	fmt.Println(personita.correo)
	personita2 := Persona{"Awita", 69, "Wembiedecoco@hotmail.com"}
	fmt.Println(personita2)
	fmt.Println(personita2.correo)

}