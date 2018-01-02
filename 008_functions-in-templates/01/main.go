package main

import (
	"html/template"
	"log"
	"os"
	"strings"
)

var myTMPL *template.Template

type person struct {
	Name string
}

//сначала создадим карту функ-й
// "uc" - фун-я которую мы вызываем в шаблоне
// "uc" - ф-я из пакета strings
// "ft" - наша объявленняая функ-я
// "ft" - возвращает первые 3 знач-я из slice[]string

var funcMap = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
}

func init() {
	//вернет ошибку, при парсинге нашего шаблона uc и ft не опеределены в нем, это статич файл
	//myTMPL = template.Must(template.ParseFiles("tmpl.gohmtl")
	//myTMPL = myTMPL.Funcs(funcMap)

	//получаем указатель на шаблон
	//template.New получает строку, и указатель на шаблон, как только у нас есть указатель на шаблон
	//можем вызывать фукнции Funcs и остальные
	myTMPL = template.Must(template.New("").Funcs(funcMap).ParseFiles("tmpl.gohtml"))
}

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	s = s[:3]
	return s
}

func main() {
	a := person{
		Name: "AmandaPrel",
	}

	b := person{
		Name: "Karlos",
	}

	c := person{
		Name: "Elliot",
	}

	peopleSlice := []person{a, b, c}

	data := struct {
		Persons []person
	}{
		peopleSlice,
	}

	myHTML := myTMPL.ExecuteTemplate(os.Stdout, "tmpl.gohtml", data)
	if myTMPL != nil {
		log.Fatalln(myHTML)
	}
}
