package main

import (
	"net/http"
	"io"
)

func main() {
	http.HandleFunc("/", test)
	//handle takes a handler
	http.Handle("/resources/", http.StripPrefix("/resources", http.FileServer(http.Dir("./assets"))))
	http.ListenAndServe(":8080", nil)
}

func test(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8)")
	io.WriteString(w, `<img src="/resources/test.jpg"`)
}
