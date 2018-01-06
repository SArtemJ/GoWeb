package main

import (
	"html/template"
	"log"
	"os"
)

var myTMPL *template.Template

func init() {
	myTMPL = template.Must(template.ParseFiles("tmpl.gohtml"))
}

func main() {
	mySlice := []string{"one", "two", "three", "zero"}

	myHTML := myTMPL.Execute(os.Stdout, mySlice)
	if myHTML != nil {
		log.Fatalln(myHTML)
	}
}
