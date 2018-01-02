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

	userTwo := person{
		Name: "UserTwo",
		Age:  40,
	}

	userThree := person{
		Name: "UserThree",
		Age:  25,
	}

	userSlice := []person{userOne, userTwo, userThree}

	myHTML := myTMPL.Execute(os.Stdout, userSlice)
	if myHTML != nil {
		log.Fatalln(myHTML)
	}
}
