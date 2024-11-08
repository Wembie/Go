package main

import (
	"fmt"
	"golang.org/x/exp/constraints"
)

/*func PrintList(list ...interface{}) {
	for _, value := range list {
		fmt.Println(value)
	}
}*/

func PrintList(list ...any) {
	for _, value := range list {
		fmt.Println(value)
	}
}

func Sum(nums ...int) (total int) {
	for _, num := range nums {
		total += num
	}
	return total
}

type integer int

/*func SumMela[T int | float64 | integer](nums ...T) (total T) {
	for _, num := range nums {
		total += num
	}
	return total
}*/

type Numbers interface{
	~int | ~float64 | ~float32 | ~uint
}

//Restriccion de aproximacion y ya serviria tambien el tipo de dato "integer"
func SumMela[T ~int | ~float64](nums ...T) (total T) {
	for _, num := range nums {
		total += num
	}
	return total
}

func SumR[T Numbers](nums ...T) (total T) {
	for _, num := range nums {
		total += num
	}
	return total
}

func SumMOD[T constraints.Integer | constraints.Float](nums ...T) (total T) {
	for _, num := range nums {
		total += num
	}
	return total
}

//Utilizar operadores de igualdad o distinto
func Includes[T comparable](list []T, value T) bool {
	for _, item := range list {
		if item == value{
			return true
		}
	}
	return false
}

func Filter[T constraints.Ordered](list []T, callback func(T) bool) []T{
	result := make([]T, 0, len(list))
	for _, item := range list {
		if callback(item){
			result = append(result, item)
		}

	}
	return result
}

/*type Product struct{
	Id uint
	Desc string
	Price float32
}*/

type Product[T uint | string] struct{
	Id T
	Desc string
	Price float32
}

func main() {
	PrintList("Juan", "Acosta", "Ana", "Fercho")
	PrintList(69, 96, 21, 12)
	PrintList(68,"+1", 69, true, 3.1416)
	fmt.Println(Sum(1,2,3,4,5)) //15
	//fmt.Println(Sum(1,2,3,4,5,5.4)) //Error porq ue arriba seria int
	fmt.Println(SumMela(1,2,3,4,5)) //15
	fmt.Println(SumMela(1,2,3,4,5,5.4)) //20.4
	var num1, num2 integer = 100, 300
	fmt.Println(SumMela(num1,num2))
	//
	fmt.Println(SumR(1,2,3,4,5)) //15
	fmt.Println(SumR(1,2,3,4,5,5.4)) //20.4
	fmt.Println(SumR(num1,num2))
	//
	fmt.Println(SumMOD[uint](1,2,3,4,5)) //15
	fmt.Println(SumMOD(1,2,3,4,5,5.4)) //20.4
	fmt.Println(SumMOD(num1,num2))

	strings := []string{"a", "b", "c", "d"}
	numbers := []int{1, 2, 3, 4}
	fmt.Println(Includes(strings, "a"))
	fmt.Println(Includes(strings, "e"))
	fmt.Println(Includes(numbers, 4))
	fmt.Println(Includes(numbers, 5))

	fmt.Println(Filter(numbers, func(value int ) bool {return value > 3}))
	fmt.Println(Filter(strings, func(value string ) bool {return value > "b"}))

	product1 := Product[uint]{1, "Zapatos", 50}
	product2 := Product[string]{"324D-232F-4234-23FS", "Zapatos", 50}
	fmt.Println(product1)
	fmt.Println(product2)
}