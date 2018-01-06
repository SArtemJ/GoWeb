package main

import (
	"html/template"
	"os"
	"log"
)

type hotel struct {
	Name string
	Address string
	City string
	Zip string
	Region string
}

type region struct {
	Region string
	Hotels []hotel
}

type Regions []region

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseFiles("tmpl.gohtml"))
}

func main() {
	hotel1 := hotel{
		Name: "First",
		Address: "Address first",
		City: "City first",
		Zip: "Zip first",
		Region: "Region first",
	}

	hotel2 := hotel{
		Name: "Two",
		Address: "Address two",
		City: "City two",
		Zip: "Zip two",
		Region: "Region two",
	}

	hotel3 := hotel{
		Name: "Three",
		Address: "Address three",
		City: "City three",
		Zip: "Zip three",
		Region: "Region three",
	}

	reg1 := region{
		Region: "Reg One",
		Hotels:[]hotel{
			hotel1,
			hotel2,
		},
	}

	reg2 := region{
		Region: "Reg Two",
		Hotels:[]hotel{
			hotel3,
		},
	}

	regions := Regions{reg1, reg2}

	err := tmpl.Execute(os.Stdout, regions)
	if err != nil{
		log.Fatalln(err)
	}


}

