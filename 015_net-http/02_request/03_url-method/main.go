package main

import (
	"html/template"
	"log"
	"net/http"
	"net/url"
)

type myHandler int

var tmpl *template.Template

func (mh myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
	}

	data := struct {
		Method      string
		URL         *url.URL
		Submissions url.Values
	}{
		r.Method,
		r.URL,
		r.Form,
	}

	tmpl.ExecuteTemplate(w, "index.gohtml", data)
}

func init() {
	tmpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	var myH myHandler
	http.ListenAndServe(":8080", myH)
}
