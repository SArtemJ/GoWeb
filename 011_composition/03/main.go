package main

import (
	"html/template"
	"os"
	"log"
)

var tmpl * template.Template

type user struct {
	Name string
	SecondName string
}

type lessons struct {
	Lessons string
	Users []user
}
type schedule struct {
	Month string
	Lessons lessons
}

func init() {
	tmpl = template.Must(template.ParseFiles("tmpl.gohtml"))
}

func main() {
	j1 := schedule{
		Month: "January",
		Lessons: lessons{
			Lessons: "Mathematic",
			Users: []user {
				user{"One", "One"},
				user{"Two", "Two"},
				},
		},
	}

	err := tmpl.Execute(os.Stdout, j1)
	if err != nil {
		log.Fatalln(err)
	}
}