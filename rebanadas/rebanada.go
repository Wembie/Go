package main
//SLICE
import "fmt"

func main() {
	//dinamicas
	dias_semana := []string{"D","L","M","M","J","V","S"}
	dia_rebanada := dias_semana[1:6]
	fmt.Println(dias_semana)
	fmt.Println(dia_rebanada)
	fmt.Println(len(dia_rebanada))//tamaño
	fmt.Println(cap(dia_rebanada))//capicidad
	dia_rebanada = append(dia_rebanada, "V") //append devuelve la misma lista
	fmt.Println(dia_rebanada)
	fmt.Println(len(dia_rebanada))//tamaño
	fmt.Println(cap(dia_rebanada))//capicidad
	dia_rebanada = append(dia_rebanada[:2],dia_rebanada[3:]... ) //[0:2]
	fmt.Println(dia_rebanada)
	fmt.Println(len(dia_rebanada))//tamaño
	fmt.Println(cap(dia_rebanada))//capicidad
	//make
	nombres := make([]string, 5, 10) //tipo, cantidad, capacidadMax
	nombres[0] = "Juan"
	fmt.Println(nombres)
	//copy
	slice := []int{0,1,2,3,4,5,6,7,8,9,}
	slice2 := make([]int, 10)
	how_many_copies_did := copy(slice2,slice)
	fmt.Println(slice)
	fmt.Println(slice2)
	fmt.Println(how_many_copies_did)

}