package main

import (
	"net/http"
	"io"
)

func example(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html; charset=UTF-8")

	io.WriteString(w, `
	<!--image doesn't serve'-->
	<img src="/test.jpg">
	`)
}

func main() {
	http.HandleFunc("/", example)
	http.ListenAndServe(":8080", nil)
}
