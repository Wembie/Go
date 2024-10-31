package main

import (
	"fmt"
	"log"

	"github.com/Wembie/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0) //bandera de formado, para no incluir fecha y hora

	names := []string{"Wumpux","Toe","BO6", "Wembie"}
	messages, err := greetings.Hellos(names)
	if err != nil{
		log.Fatal(err)
	}
	/*message, err := greetings.Hello("Wembie")
	//message, err := greetings.Hello("")
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println(message)*/
	fmt.Println(messages)
}