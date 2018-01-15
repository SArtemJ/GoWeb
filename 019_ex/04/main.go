package main

import (
	"net/http"
	"html/template"
	"log"
)

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseFiles("templates/index.gohtml"))
}

func main() {
	http.HandleFunc("/", test)
	http.Handle("/resources/", http.StripPrefix("/resources", http.FileServer(http.Dir("./imgs"))))
	http.ListenAndServe(":8080", nil)
}

func test(w http.ResponseWriter, r *http.Request) {
	err := tmpl.Execute(w, nil)
		if err != nil {
			log.Fatalln("template not execute", err)
		}
}