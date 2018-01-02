package main

import (
	"html/template"
	"log"
	"os"
)

var myTMPL *template.Template

type person struct {
	Name string
	Age  int
}

func init() {
	myTMPL = template.Must(template.ParseFiles("tmpl.gohtml"))
}

func main() {
	userOne := person{
		Name: "UserOne",
		Age:  32,
	}

	myHTML := myTMPL.Execute(os.Stdout, userOne)
	if myHTML != nil {
		log.Fatalln(myHTML)
	}
}
