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

	g1 := struct {
		Score1 int
		Score2 int
	}{
		7,
		9,
	}

	myHTML := myTMPL.Execute(os.Stdout, g1)
	if myHTML != nil {
		log.Fatalln(myHTML)
	}
}
