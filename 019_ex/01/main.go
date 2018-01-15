package main

import (
	"net/http"
	"io"
	"html/template"
	"log"
)

/*
ListenAndServe on port 8080 of localhost
For the default route "/" Have a func called "foo" which writes to the response "foo ran"

For the route "/dog/"
Have a func called "dog" which parses a template called "dog.gohtml" and writes to the response
"This is from dog"
and also shows a picture of a dog when the template is executed.
Use "http.ServeFile" to serve the file "dog.jpeg"

*/

func main() {

	http.HandleFunc("/", blankPage)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/test.jpg", showIMG)
	http.ListenAndServe(":8080", nil)
}


func blankPage(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "foo ran")
}

func dog(w http.ResponseWriter, r *http.Request) {
	page, err := template.ParseFiles("image.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	page.ExecuteTemplate(w,"image.gohtml", nil)

}

func showIMG(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "test.jpg")
}
