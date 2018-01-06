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

	data := struct {
		Words []string
		Lname string
	}{
		mySlice,
		"UserName",
	}

	myHTML := myTMPL.Execute(os.Stdout, data)
	if myHTML != nil {
		log.Fatalln(myHTML)
	}
}
