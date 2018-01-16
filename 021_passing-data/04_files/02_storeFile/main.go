package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"os"
	"html/template"
)

var tmpl *template.Template

func init()  {
	tmpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", funcT)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func funcT(w http.ResponseWriter, r *http.Request) {

	var s string
	fmt.Println(r.Method)
	if r.Method == http.MethodPost {
		//open our file
		file, header, err := r.FormFile("some")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer file.Close()

		fmt.Println("\nfile: ", file, "\nheader", header, "\nerr", err)

		//read from file
		bs, err := ioutil.ReadAll(file)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		s = string(bs)

		dst, err := os.Create(filepath.Join("./userFolder/", header.Filename))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer dst.Close()

		_, err = dst.Write(bs)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}

	w.Header().Set("Content-type", "text/html; charset=utf-8")
	tmpl.ExecuteTemplate(w, "index.gohtml", s)
}
