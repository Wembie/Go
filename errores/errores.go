package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
)

func div( diviendo, divisor int )( int, error){
	if divisor == 0{
		return 0, errors.New("No se puede divir por cero pa, cha, loka")
		//return 0, fmt.Errorf("No se puede divir por cero pa, cha, loka")
	}
	return diviendo / divisor, nil
}

func dividir( diviendo, divisor int ){
	fmt.Println(diviendo / divisor) 
}

func divide( diviendo, divisor int ){
	//Funcion anonima
	defer func ()  {
		if r := recover(); r != nil{
			fmt.Println(r)
		}
	}()
	validate_zero(divisor)
	fmt.Println(diviendo / divisor) 
}


func validate_zero(divisor int){
	if divisor == 0{
		panic("No se puede divir por cero pa, cha, loka")
	}
}

func main() {
	str := "123"
	//str := "123f"
	num, err := strconv.Atoi(str)
	//En Go, nil es el valor cero para punteros, interfaces, mapas, slices, canales y funciones; y corresponde a la representación de un valor no inicializado.
	if err != nil{
		fmt.Println("Error: ", err)
		return
	}
	fmt.Println("Numero: ", num)
	//result, er := div(10,0)
	result, er := div(10,2)
	if er != nil{
		fmt.Println("Error: ", er)
		return
	}
	fmt.Println("Resultado: ", result)

	//Defer para posponer la ejecucion de una funcion hasta que haya finalizado
	
	fmt.Println(3) //3
	fmt.Println(1) //1
	fmt.Println(2) //2

	defer fmt.Println(3) //Imprimiria a lo ultimo DE LA FUNCION JAJA
	fmt.Println(1) //1
	fmt.Println(2) //2
	//y si hay varios seria una fila
	defer fmt.Println(3) //1
	defer fmt.Println(2) //2
	defer fmt.Println(1) //3

	/*file, errorcito := os.Create("Holita.txt")
	if errorcito != nil{
		fmt.Println("Error: ", errorcito)
		return
	}
	_, errorcito = file.Write([]byte("Hola, Wumpux"))
	if errorcito != nil{
		fmt.Println("Error: ", errorcito)
		file.Close() //PARA NO REPETIR TANTO CODIGO SE USARIA DEFER
		return
	}
	file.Close()*/

	file, errorcito := os.Create("Holita.txt")
	if errorcito != nil{
		fmt.Println("Error: ", errorcito)
		return
	}
	defer file.Close()
	_, errorcito = file.Write([]byte("Hola, Wumpux"))
	if errorcito != nil{
		fmt.Println("Error: ", errorcito)
		return
	}

	//Panic y recover
	divide(100,10)
	divide(200,25)
	//divide(100,0) //Se paniquea xd y muere
	divide(100,0) //Ya que se usa el recover entonces continua el programa aunque se paniquee xd
	//dividir(100,0) //Igual se paniquea por que tiene el panic automaitoc del panic: runtime error: integer divide by zero
	divide(69,9)

	//Registro de errores LOGS
	
	log.Print("Mensaje de registro")
	log.Println("Otro mensaje de registro")
	//log.Fatal("REGISTRO DE ERRORRRRRRRRRRRRRRRRRRRRRR") //Muere el programa
	//log.Panic("Se paniquio") //Tambien muere

	log.SetPrefix("main(): ")
	log.Print("Estoy")
	{
	file, err := os.OpenFile("info.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644) //crearlo si existe si no lo abre| esacribir y abrir solo lectura y el codigo es para WyR
	if err != nil{
		log.Fatal(err)
	}
	defer file.Close()
	
	log.SetOutput(file)
	log.Print("Estamos activos o no tamos activos")

	}
}
  