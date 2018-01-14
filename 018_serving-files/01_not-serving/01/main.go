package main

import (
	"net/http"
	"io"
)

func example(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html; charset=UTF-8")

	io.WriteString(w, `
	<!--not serving from our server-->
	<img src="https://upload.wikimedia.org/wikipedia/commons/6/6e/Golde33443.jpg">
	`)
}

func main() {
	http.HandleFunc("/", example)
	http.ListenAndServe(":8080", nil)
}
