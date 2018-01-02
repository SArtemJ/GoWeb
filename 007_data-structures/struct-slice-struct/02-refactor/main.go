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

type pet struct {
	TypePET string
}

//type items struct {
//	Users []person
//	Pets []pet
//}

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

	petOne := pet{
		TypePET: "Cat",
	}

	petTwo := pet{
		TypePET: "Dog",
	}

	userSlice := []person{userOne, userTwo, userThree}
	petSlice := []pet{petOne, petTwo}

	itemsValue := struct{
		Users []person
		Pets []pet
		}{
		userSlice,
		petSlice,
	}

	myHTML := myTMPL.Execute(os.Stdout, itemsValue)
	if myHTML != nil {
		log.Fatalln(myHTML)
	}
}
