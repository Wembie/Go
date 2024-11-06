package book

import "fmt"

//Polimorfismo

type Printable interface{
	PrintInfo()
}

func Print(p Printable){
	p.PrintInfo()
}

//SI SE HACEN LA PRIMERA LETRA MINUSCULA SERIAN PRIVADOS Y EN MAYUSCULA LA PRIMERA SERIA PUBLICA OSEA ENCAPSULAMIENTO TANTO PARA METODOS COMO ATRIBUTOS

type Book struct {
	/*Title  string
	Author string
	Pages  int*/
	title  string
	author string
	pages  int
}

//No existe constructor pero se puede simular

func NewBook( title, author string, pages int) *Book {
	return &Book{
		title: title,
		author: author,
		pages: pages,
	}
}

func (b *Book) SetTitle( title string){
	b.title = title
}

func (b *Book) GetTitle() string{
	return b.title
}

func (b *Book) PrintInfo() {
	fmt.Printf("Title: %s\nAuthor: %s\nPages: %d\n\n", b.title, b.author, b.pages)
}

//herencia
type TextBook struct {
	Book
	editorial string
	lvl string
}

func NewTextBook( title, author string, pages int, editorial, lvl string) *TextBook {
	return &TextBook{
		Book: Book{title, author,pages},
		editorial: editorial,
		lvl: lvl,
	}
}

func (b *TextBook) PrintInfo() {
	fmt.Printf("Title: %s\nAuthor: %s\nPages: %d\nEditorial: %s\nLvL: %s\n\n", b.title, b.author, b.pages, b.editorial, b.lvl)
}