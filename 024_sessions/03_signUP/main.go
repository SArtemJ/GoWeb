package main

import (
	"html/template"
	"net/http"
)

type user struct {
	UserMail string
	Fname string
	Lname string
}

var tmpl *template.Template
var dbUsers = map[string]string{}
var dbSessions = map[string]user{}

func init() {
	tmpl = template.Must(template.ParseGlob("templates/*"))
}

func main()  {
	http.HandleFunc("/", indexPage)
	http.HandleFunc("/signup", signup)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func indexPage(w http.ResponseWriter, r *http.Request) {

}