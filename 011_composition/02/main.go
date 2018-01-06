package main

import (
	"html/template"
	"log"
	"os"
)

type person struct {
	Name string
	Age int
}

type doubleZero struct {
	person
	DriveLicense bool
}

var myTMPL *template.Template

func init() {
	myTMPL = template.Must(template.ParseFiles("tmpl.gohtml"))
}

func main() {
	p1 := doubleZero{
		person{
			Name: "Elliot Alderson",
			Age: 27,
		},
		true,
	}

	myHTML := myTMPL.Execute(os.Stdout, p1)
	if myHTML != nil {
		log.Fatalln(myHTML)
	}
}
