package main

import (
	"html/template"
	"os"
	"log"
)

var myTmpl *template.Template

func init() {
	myTmpl = template.Must(template.ParseFiles("tmpl.gohtml"))
}

func main() {

	mySlice := []string{"One", "Two", "Zero"}

	myHTML := myTmpl.Execute(os.Stdout, mySlice)
	if myHTML != nil {
		log.Fatalln(myHTML)
	}

}
