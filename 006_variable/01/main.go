package main

import (
	"html/template"
	"log"
	"os"
)

var myTpl *template.Template

func init() {
	myTpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	resultHTML := myTpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", ` ;345; `)
	if resultHTML != nil {
		log.Fatalln(resultHTML)
	}
}
