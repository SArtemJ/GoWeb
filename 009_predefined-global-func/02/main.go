package main

import (
	"html/template"
	"log"
	"os"
)

var myTMPL *template.Template

type user struct {
	Name string
	SecondName string
	Admin bool
}

func init() {
	myTMPL = template.Must(template.ParseFiles("tmpl.gohtml"))
}

func main() {

	u1 := user{
		Name: "Bill",
		SecondName: "Joy",
		Admin: false,
	}

	u2 := user{
		Name: "",
		SecondName: "Admin",
		Admin: true,
	}

	u3 := user{
		Name: "Sara",
		SecondName: "Conor",
		Admin: true,
	}

	users := []user{u1, u2, u3}

	myHTML := myTMPL.Execute(os.Stdout, users)
	if myHTML != nil {
		log.Fatalln(myHTML)
	}
}
