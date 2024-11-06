package main

import "fmt"

//Funcion variadica osea que tiene N parametros
func suma(name string, nums ...int)(total int){
	//fmt.Printf("%T - %v\n", nums, nums)
	for _, num := range nums {
		total += num
	}
	fmt.Printf("Que se dice el/la %s, acordate que la suma de esos numeros que mando fueron de: ", name)
	return

}
//Varios parametros de varios tipos de datos
func imprimirDatos(datos ...interface{}){
	for _, dato := range datos{
		fmt.Printf("%T - %v\n", dato, dato)
	}
}

//RECURSIVIDADDDDDDDDD

func factorial(n int) int {
	if n == 0{
		return 1
	}
	return n * factorial(n - 1)
}

//Funciones como valores
func saludar(name string, f func(string)){
	f(name)
}

func duplicar(n int) int {
	return n * 2
}

func triplicar(n int) int {
	return n * 3
}

func aplicar(f func(int) int, n int) int {
	return f(n)
}


//funcion orden superior
func double( f func(int) int, x int) int{
	return f(x * 2)
}

func addOne( x int ) int{
	return x + 1
}

//Clousures
//retorna otra funcion
func incrementer() func() int {
	i := 0
	return func() int{
		i++
		return i
	}
}

func main() {
	//fmt.Println(suma(12,45,78,56))
	//fmt.Println(suma(1,2,3,4,5,6,7,8,9))
	fmt.Println(suma("Wumpux",1,2,3,4,5,6,7,8,9))
	fmt.Println(suma("Wumpux",1,68))
	imprimirDatos("Wembie",69,true,69.69)

	fmt.Println("Factorial:", factorial(3))

	//FUNCION ANONIMA (SIN NOMBRE)
	func (){
		fmt.Printf("Soy anonymous pa\n")
	}() //toca o si no le coloca parametros

	saludo := func (name string)  {
		fmt.Printf("Que se dice %s! Bien o q?\n", name)
	}
	saludo("Isa")
	saludar("Wembiesito", saludo)

	r1 := aplicar(duplicar, 5)
	r2 := aplicar(triplicar, 5)
	fmt.Println(r1, r2)

	r := double(addOne, 3)
	fmt.Println(r)
	
	//Clousures
	//Los closures permiten que una función acceda y mantenga referencias a variables fuera de su propio alcance léxico.
	nextInt := incrementer()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
}	