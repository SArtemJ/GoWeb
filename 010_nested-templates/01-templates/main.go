package main

import (
	"html/template"
	"os"
	"log"
)

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseGlob("templates/*"))
}

func main()  {
	err := tmpl.ExecuteTemplate(os.Stdout, "index.gohtml", "special")
	if err != nil {
		log.Fatalln(err)
	}
}