package main

import (
	"fmt"
	"math"
	"strconv"
)


func main() {
	//int8 = -128 | 127, etc etc etc (poner mouse en el type para ver el rango)
	var integer int8 = 127
	var positivo uint8 = 255
	var float float32 = 69.69
	fmt.Println(integer,float,positivo)
	fmt.Println(math.MinInt64, math.MaxInt64, math.MaxFloat64, math.MaxUint8)
	var valueBoolT, valueBoolF bool = true, false

	fullname := "Juan Acosta \t(alias \"Wembie\")\n"
	fmt.Println(valueBoolT, valueBoolF, fullname)

	var a byte = 'a'
	fmt.Println(a) //ASCII

	s := "Awita"
	fmt.Println(s[0])

	var r rune = '😻'
	fmt.Println(r)

	//Valores predeterminados

	var (
		defaultInt int //0
		defaultUint uint //0
		defaultFloat float32 //0
		defaultBool bool //false
		defaultString string //" "
	)
	fmt.Println(defaultInt, defaultUint, defaultFloat, defaultBool, defaultString)

	//Conversiones
	var integer16 int16 = 50
	var integer32 int32 = 100
	//fmt.Println(integer16 + integer32) MALA
	fmt.Println(integer16 + int16(integer32))

	f := "100"
	i, _ := strconv.Atoi(f) //Convertir str to int
	fmt.Println(i)
	fmt.Println(i+69)

	n := 42
	f = strconv.Itoa(n) //Convertir int to str
	fmt.Println(f)
	fmt.Println(f+f)
}