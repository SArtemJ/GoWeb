package main

import (
	"html/template"
	"os"
	"log"
)

type person struct {
	Name string
	Age int
}

func (p person) SomeProcessing() int {
	return 20
}

func (p person) AgeDbl() int {
	return p.Age * 2
}

func (p person) TakesAge(x int) int {
	return x * 2
}

var tmpl * template.Template

func init()  {
	tmpl = template.Must(template.ParseFiles("tmpl.gohtml"))
}

func main() {
	p := person{
		Name: "Elliot",
		Age: 27,
	}

	err := tmpl.Execute(os.Stdout, p)
	if err != nil {
		log.Fatalln(err)
	}
}


//best practice: call func's im templates fo formatting only; not processing or data access