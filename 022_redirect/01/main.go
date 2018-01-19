package main

import (
	"html/template"
	"net/http"
	"fmt"
)

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {

	http.HandleFunc("/", blankPage)
	http.HandleFunc("/star", star)
	http.HandleFunc("/stara", stara)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)

}

func blankPage(w http.ResponseWriter, r *http.Request) {
	fmt.Print("Request method is ", r.Method, "\n\n")
}

func star(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request method is ", r.Method)
	w.Header().Set("Location", "/")
	w.WriteHeader(http.StatusSeeOther)
}

func  stara(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("Request is ", r.Method)
	tmpl.ExecuteTemplate(w, "index.gohtml", nil)
}