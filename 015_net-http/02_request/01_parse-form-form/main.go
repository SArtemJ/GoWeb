package main

import (
	"html/template"
	"log"
	"net/http"
)

type myHandler int

var tmpl *template.Template

func (mh myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}

	tmpl.ExecuteTemplate(w, "index.gohtml", r.Form)
}

func init() {
	tmpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	var myH myHandler
	http.ListenAndServe(":8080", myH)
}
