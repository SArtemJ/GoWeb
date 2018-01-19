package main

import (
	"html/template"
	"net/http"
	"fmt"
)

var tmpl *template.Template

func init()  {
	tmpl = template.Must(template.ParseGlob("templates/*"))
}

func main()  {

	http.HandleFunc("/", blankPage)
	http.HandleFunc("/star", star)
	http.HandleFunc("/starb", starb)
	http.ListenAndServe(":8080", nil)

}

func blankPage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request method is ", r.Method)
}

func star(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request method is ", r.Method)
	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}

func starb (w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request method is", r.Method)
	tmpl.ExecuteTemplate(w, "index.gohtml", nil)
}