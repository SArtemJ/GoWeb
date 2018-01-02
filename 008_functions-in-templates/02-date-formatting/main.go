package main

import (
	"html/template"
	"log"
	"os"
	"time"
)

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.New("").Funcs(fm).ParseFiles("tmpl.gohtml"))
}

func monthDayYear(t time.Time) string {
	return t.Format("02-01-2018")
}

var fm = template.FuncMap{
	"fdateMDY": monthDayYear,
}

func main() {
	html := tmpl.ExecuteTemplate(os.Stdout, "tmpl.gohtml", time.Now())
	if html != nil {
		log.Fatalln(html)
	}
}
