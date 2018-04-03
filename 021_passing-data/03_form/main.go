package main

import (
	"net/http"
	"html/template"
	"log"
)

var tmpl *template.Template

func init()  {
	tmpl = template.Must(template.ParseGlob("templates/*"))
}

type person struct {
	FirstName string
	LastName string
	Subscribed bool
}

func main() {
	http.HandleFunc("/", funcT)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func funcT(w http.ResponseWriter, r *http.Request) {

	fName := r.FormValue("firstName")
	log.Println(fName)
	sName := r.FormValue("secondName")
	log.Println(sName)
	sV := r.FormValue("subscribeValue") == "on"
	log.Println(sV)

	err := tmpl.ExecuteTemplate(w, "index.gohtml", person{fName, sName, sV})
	if err != nil {
		http.Error(w, err.Error(), 500)
		log.Fatalln(err)
	}
}
