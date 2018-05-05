package main

import (
	"net/http"
	"log"
	"html/template"
)

type person struct {
	Lang string
	L string
	Ms string
}


func main() {



	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))


}


func handler(w http.ResponseWriter, r *http.Request) {

	p := person{
		Lang: "en",
		Ms: "fuck",
	}


		t, err := template.ParseFiles("tmpl.tmpl")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = t.Execute(w, p)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

}