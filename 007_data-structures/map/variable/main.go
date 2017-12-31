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

	myMAP := map[int]string{
		0: "Alis",
		1: "Bob",
		3: "Tom",
		4: "Elliot"}

	myHTML := myTMPL.Execute(os.Stdout, myMAP)
	if myHTML != nil {
		log.Fatalln(myHTML)
	}
}
