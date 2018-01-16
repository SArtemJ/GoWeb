package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"io"
)

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
	}

	w.Header().Set("Content-type", "text/html; charset=utf-8")
	io.WriteString(w, `<form method="POST" enctype="multipart/form-data">
	<input type="file" name="some">
	<input type="submit">
	</form>
	<br>`+s)
}
