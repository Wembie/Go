package main

import (
	//"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"regexp"
)

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	//0600 lectura y escritura
	return os.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string)(*Page, error){
	filename := title + ".txt"
	body, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")
/*func getTitle(w http.ResponseWriter, r *http.Request) (string, error) {
    m := validPath.FindStringSubmatch(r.URL.Path)
    if m == nil {
        http.NotFound(w, r)
        return "", errors.New("Título de página inválido")
    }
    return m[2], nil // El título es la segunda subexpresión.
}*/

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2])
	}
}



/*func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola me encantas los %s", r.URL.Path[1:])
}*/
func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	//title := r.URL.Path[len("/view/"):]
	/*title, err := getTitle(w, r)
	if err != nil{
		return
	}*/
	p, err := loadPage(title)
	if err != nil{
		http.Redirect(w, r,"/edit/"+title, http.StatusFound)
		return
	}
	//fmt.Fprintf(w, "<h1>%s<h1> <div>%s<div>", p.Title, p.Body)
	/*t, _ := template.ParseFiles("view.html")
	t.Execute(w,p)*/
	renderTemplates(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request, title string) {
	//title := r.URL.Path[len("/edit/"):]
	/*title, err := getTitle(w, r)
	if err != nil{
		return
	}*/
	p, err := loadPage(title)
	if err != nil{
		p = &Page{Title: title}
	}
	//cargar paguina html
	/*t, _ := template.ParseFiles("edit.html")
	t.Execute(w,p)*/
	renderTemplates(w, "edit", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
	//title := r.URL.Path[len("/save/"):]
	/*title, err := getTitle(w, r)
	if err != nil{
		return
	}*/
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	err := p.save()
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

var templates = template.Must(template.ParseFiles("edit.html", "view.html"))

func renderTemplates(w http.ResponseWriter, tmpl string, p *Page){
	/*t, err := template.ParseFiles(tmpl + ".html")
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}*/
	
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}


func main() {
	fmt.Println("WEB pa")
	/*p1 := &Page{Title: "TestPage", Body: []byte("Paguina de muestra")}
	p1.save()

	p2, _ := loadPage("TestPage")
	fmt.Println(string(p2.Body))*/
	//http.ListenAndServe("localhost:8080", nil) //Es lo mismo

	/*http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hola me encantas los %s", "Tiburones")
	})*/
	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))
	log.Fatal(http.ListenAndServe(":8080", nil))
}