package main

import (
	"fmt"
	"log"
	"net/http"
	"html/template"
)


var templates = template.Must(template.ParseFiles("./templates/view.html"))

func renderTemplate(w http.ResponseWriter, tmpl string) {
	err := templates.Execute(w, tmpl+".html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func makeHandler(fn func (http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
        fn(w, r, r.URL.Path)
	}
}


func indexHandler(w http.ResponseWriter, r *http.Request, title string) {
	switch r.Method {
	case "GET":
		renderTemplate(w, "view")
	default:
		
	} 
	
}


func main() {
	
	fmt.Println("Server starting...")

	http.HandleFunc("/", makeHandler(indexHandler))
	log.Fatal(http.ListenAndServe(":8080", nil))
}