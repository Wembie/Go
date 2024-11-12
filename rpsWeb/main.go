package main

//go install github.com/air-verse/air@latest
//Para usarlo que es como el --reload en python
//comando: "air init"
//y para correrlo seria con: "air"
//y para cerrarlo seria con: "ctrl + c"

import (
	"fmt"
	"log"
	"net/http"
	"rpsWeb/handlers"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Que se dice")
}

func main() {
	//Crear enrutador
	router := http.NewServeMux()

	//Archivos estaticos
	fs := http.FileServer(http.Dir("static"))
	//ruta para acceder a archivos estaticos
	router.Handle("/static/", http.StripPrefix("/static/", fs))

	router.HandleFunc("/", handlers.Index)
	router.HandleFunc("/new", handlers.NewGame)
	router.HandleFunc("/game", handlers.Game)
	router.HandleFunc("/play", handlers.Play)
	router.HandleFunc("/about", handlers.About)

	port := ":8080"
	log.Printf("Escuchando http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, router))
}