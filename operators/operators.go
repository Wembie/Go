package main

import (
	"fmt"
	"math"
)

func main() {
	x := 69
	y := 96
	fmt.Println(x + y)
	fmt.Println(x - y)
	fmt.Println(x * y)
	fmt.Println(x / y)
	fmt.Println(x % y)

	//Incrementables
	x++ //69+1
	//Decrementables
	y-- //96-1
	//x = x - 96
	x -= 96 // x=69+1-96
	fmt.Println(x + y)
	fmt.Println(math.Pi)
	fmt.Println(math.E)
	pow := math.Pow(2,4)
	fmt.Println(pow)
	fmt.Println(math.Sqrt(64))
	fmt.Println(math.Cbrt(27)) //cubica

	// Comparación de números
    fmt.Println(1 == 1)   // true
    fmt.Println(1 != 2)   // true
    fmt.Println(2 < 3)    // true
    fmt.Println(3 > 4)    // false
    fmt.Println(4 <= 4)   // true
    fmt.Println(5 >= 6)   // false
 
    // Comparación de cadenas
    fmt.Println("hola" == "hola")       // true
    fmt.Println("hola" != "adios")      // true
    fmt.Println("abc" < "def")          // true
    fmt.Println("ghi" > "fgh")          // true
    fmt.Println("hij" <= "hij")         // true
    fmt.Println("klm" >= "klmno")       // false
 
    // Comparación de booleanos
    fmt.Println(true == true)           // true
    fmt.Println(false != true)          // true
    fmt.Println(true && false == false) // true
    fmt.Println(true || false == true)  // true

	a := true
	b := false
	c := a && b //AND
	fmt.Println(c) // Imprime false

	d := true
	e := false
	f := d || e //OR
	fmt.Println(f) // Imprime false

	g := true
	h := !g //NOT
	fmt.Println(h) // Imprime false

	// Si queremos usar `x` como bool en un nuevo contexto
	//NUEVO BLOQUE
	{
		x := true
		y := false
 
    // Negación
    fmt.Println(!x) // false
    fmt.Println(!y) // true
 
    // AND lógico
    fmt.Println(x && x) // true
    fmt.Println(x && y) // false
    fmt.Println(y && y) // false
 
    // OR lógico
    fmt.Println(x || x) // true
    fmt.Println(x || y) // true
    fmt.Println(y || y) // false

	}
	{
		x := 5
		y := 10
		z := 15
	
		// Expresión con paréntesis, operadores aritméticos, de comparación y lógicos
		resultado := ((x+y)*z)/(x*y) > z && x != y
	
		// Imprimir el resultado
		fmt.Println(resultado) //False
	}
}