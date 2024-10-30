package main

import "fmt"

func main() {
	colors := map[string]string{
		"rojo": "#FF0000",
		"verde": "#00FF00",
		"azul": "0000FF",
	}
	fmt.Println(colors)
	fmt.Println(colors["rojo"])
	colors["negro"] = "#000000"
	fmt.Println(colors)
	valor, ok := colors["verde"] //si existe
	fmt.Println(valor, ok)
	/*value, okey := colors["blanco"] //si no existe trae algo vacio
	if okey{
		fmt.Println("Si existe pa")
	} else{
		fmt.Println("No existe pa")
	}*/ 
	if value, okey := colors["blanco"]; okey{
		fmt.Println("Si existe pa y es:", value)
	} else{
		fmt.Println("No existe pa")
	}
	//fmt.Println(value, okey)
	delete(colors, "azul")
	fmt.Println(colors)

	for clave, valor := range colors{
		fmt.Printf("Clave: %s, Valor: %s\n", clave, valor)
	}
}