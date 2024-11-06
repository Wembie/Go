package main

import (
	"fmt"
	"poo/animal"
	"poo/book"
)

func main() {
	/*my_book := book.Book{
		Title: "Moby Dick",
		Author: "Herman Melville",
		Pages: 300,
	}*/
	my_book := book.NewBook("Moby Dick", "Herman Melville", 300)
	my_book.PrintInfo()
	my_book.SetTitle("Moby Dick (Wumpux Edition)")
	fmt.Println(my_book.GetTitle())
	my_book.PrintInfo()
	
	scholar_book := book.NewTextBook("Comunicacion", "Jaime Gamarra", 261, "Santillana SAC", "Secundaria")
	scholar_book.PrintInfo()

	//Polimorfismo

	//my_book.PrintInfo()
	//scholar_book.PrintInfo()
	fmt.Println("Polimorfismo\n")
	book.Print(my_book)
	book.Print(scholar_book)

	perro := animal.Perro{Nombre: "Toby"}
	gato := animal.Gato{Nombre: "Tom"}

	animal.HacerSonido(&perro)
	animal.HacerSonido(&gato)

	animales := []animal.Animal{
		&animal.Perro{Nombre: "Lunita"},
		&animal.Gato{Nombre: "WanWan"},
		&animal.Perro{Nombre: "Sharky"},
		&animal.Gato{Nombre: "AwitaECoco"},
	}

	for _, animal := range animales{
		animal.Sonido()
	}

}