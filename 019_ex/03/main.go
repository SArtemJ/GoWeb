package main

import (
	"html/template"
	"net/http"
	"log"
)

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseFiles("templates/index.gohtml"))
}

func main()  {
	fs := http.FileServer(http.Dir("imgs"))
	http.Handle("/simple/", fs)
	http.HandleFunc("/", myfunc)
	http.ListenAndServe(":8080", nil)
}

func myfunc(w http.ResponseWriter, r *http.Request)  {
	err := tmpl.Execute(w, nil)
	if err != nil {
		log.Fatalln("template no execute", err)
	}
}