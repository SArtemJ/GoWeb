package main

import (
	"html/template"
	"os"
	"log"
)

type food struct {
	Name string
	Price float64
}


type Menu struct {
	RestourantName string
	Food []food
}

type restor []Menu

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseFiles("tmpl.gohtml"))
}

func main() {
	food1 := food{
		Name: "Donut's",
		Price: 10.00,
	}

	food12 := food{
		Name: "Cup of tea",
		Price: 3.00,
	}

	food13 := food{
		Name: "Cup of coffee",
		Price: 4.00,
	}

	food2 := food{
		Name: "Fritters",
		Price: 7.00,
	}

	food21 := food{
		Name: "Juice",
		Price: 2.00,
	}

	food3 := food{
		Name: "Meat",
		Price: 15.00,
	}

	foodMain := []food{food1, food12, food13}
	foodSec := []food{food2, food21}
	foodLast := []food{food3}

	res := Menu{
		RestourantName: "The one",
		Food: foodMain,
	}

	res2 := Menu{
		RestourantName: "The two",
		Food: foodSec,
	}

	res3 := Menu{
		RestourantName: "The three",
		Food: foodLast,
	}

	restr := restor{res, res2, res3}

	err := tmpl.Execute(os.Stdout, restr)
	if err != nil{
		log.Fatalln(err)
	}


}

