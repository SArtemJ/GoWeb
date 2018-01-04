package main

import (
	"html/template"
	"log"
	"math"
	"os"
	"fmt"
)

var myTMPL *template.Template

func init() {
	myTMPL = template.Must(template.New("").Funcs(fm).ParseFiles("tmpl.gohtml"))

}

func double(x int) int {
	return x + x
}

func square(x int) float64 {
	return math.Pow(float64(x), 2)
}

func sqRoot(x float64) float64 {
	fmt.Println("aaa")
	fmt.Println(x)
	fmt.Println(math.Sqrt(x))
	return math.Sqrt(x)
}

var fm = template.FuncMap{
	"fdbl":    double,
	"fsquare": square,
	"fsqRoot": sqRoot,
}

func main() {

	myHTML := myTMPL.ExecuteTemplate(os.Stdout, "tmpl.gohtml", 3)
	if myTMPL != nil {
		log.Fatalln(myHTML)
	}
}
