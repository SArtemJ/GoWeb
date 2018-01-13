package main

/*Take the previous program in the previous folder and change it so that:
a template is parsed and served
you pass data into the template*/

import (
	"net/http"
	"html/template"
	"log"
)

var tmpl *template.Template

//type myText struct {
//	SimpleText string
//}

func init() {
	tmpl = template.Must(template.ParseFiles("tmpl.gohtml"))
}

func indexPage(res http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil{
		log.Fatalln(err)
	}

	tmpl.ExecuteTemplate(res, "tmpl.gohtml", "Index page")

}

func dogPage(res http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}

	tmpl.ExecuteTemplate(res, "tmpl.gohtml", "Dog page")

}

func mePage(res http.ResponseWriter, r *http.Request)  {
	err := r.ParseForm()
	if err != nil{
		log.Fatalln(err)
	}

		tmpl.ExecuteTemplate(res, "tmpl.gohtml", "Artem")
}

func main()  {

	//srvMux
	srvMux := http.NewServeMux()
	srvMux.HandleFunc("/", indexPage)
	srvMux.HandleFunc("/dog/", dogPage)
	srvMux.HandleFunc("/me/", mePage)
	http.ListenAndServe(":8080", srvMux)

	//default
	//http.HandleFunc("/", indexPage)
	//http.HandleFunc("/dog/", dogPage)
	//http.HandleFunc("/me/", mePage)
	//http.ListenAndServe(":8080", nil)

}

