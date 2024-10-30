package main

import "fmt"

type Persona struct {
	nombre string
	edad int
	correo string
}
//    RECEPTOR
func (p *Persona) diHola(){
	fmt.Println("Hola mi nombre es", p.nombre)
}

func main() {
	var x int = 10
	var p *int = &x

	fmt.Println(x)
	fmt.Println(&x)
	fmt.Println(p)
	editar(&x)
	fmt.Println(x)

	////////////

	personita := Persona{"Wembie", 22, "Awitadecoco@hotmail.com"}
	personita.diHola()
}

/*func editar( puntero int){ no cambia
	puntero = 20
}*/

func editar( puntero *int){
	*puntero = 20
}