package handlers

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"rpsWeb/rps"
	"strconv"
)

const (
	templateDir = "templates/"
	templateBase = templateDir + "base.html"
)

type Player struct{
	Name string
}

var player Player

func Index(w http.ResponseWriter, r *http.Request) {
	restarValue()
	renderTemplate(w, "index.html", nil)
	//tpl := template.Must(template.ParseFiles("templates/base.html", "templates/index.html"))
	/*tpl, err := template.ParseFiles("templates/base.html", "templates/index.html")
	if err != nil{
		http.Error(w, "Error al analizar plantilla", http.StatusInternalServerError)
		return
	}*/ //el must hace esto mismo y hace un panic en caso de error
	/*data := struct{
		Title string
		Message string
	}{
		Title: "Paguina de Inicio",
		Message: "Welcome a R, P, S",
	}
	err = tpl.Execute(w, data)
	//err = tpl.Execute(w, nil)*/

	/*err := tpl.ExecuteTemplate(w, "base", nil)
	if err != nil{
		http.Error(w, "Error al renderizar plantilla", http.StatusInternalServerError)
		return
	}*/
}

func NewGame(w http.ResponseWriter, r *http.Request) {
	restarValue()
	//fmt.Fprintln(w, "Crear nuevo juego")
	renderTemplate(w, "new-game.html", nil)
}

func Game(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintln(w, "Juego")
	if r.Method == "POST"{
		err := r.ParseForm()
		if err != nil{
			http.Error(w, "Error parsing form", http.StatusBadRequest)
			return
		}
		player.Name = r.Form.Get("name")
	}
	if player.Name == ""{
		http.Redirect(w, r, "/new", http.StatusFound)
	}
	renderTemplate(w, "game.html", player)
}

func Play(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintln(w, "Jugar")
	playerChoice, _ := strconv.Atoi(r.URL.Query().Get("c"))
	result := rps.PlayRound(playerChoice)

	//fmt.Println(result)
	out, err := json.MarshalIndent(result, "", "    ")
	if err != nil{
		log.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)

}

func About(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintln(w, "Acerca de")
	restarValue()
	renderTemplate(w, "about.html", nil)
}

func renderTemplate(w http.ResponseWriter, page string, data any){
	tpl := template.Must(template.ParseFiles(templateBase, templateDir+page))
	err := tpl.ExecuteTemplate(w, "base", data)
	if err != nil{
		http.Error(w, "Error al renderizar plantilla", http.StatusInternalServerError)
		log.Println(err)
		return
	}
}

//reiniciar valores

func restarValue(){
	player.Name = ""
	rps.ComputerScore = 0
	rps.PlayerScore = 0

}