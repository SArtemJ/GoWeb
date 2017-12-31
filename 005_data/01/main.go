package main

import (
	"html/template"
	"log"
	"os"
	"fmt"
)

var myTpl *template.Template

func init() {
	myTpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	html := myTpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", "MyPassedDataIntoTheTamplate")
	fmt.Printf("%T", html)
	if html != nil {
		log.Fatalln(html)
	}
}
