package main

import (
"net/http"
"io"
)

func main() {
	http.HandleFunc("/", funcT)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func funcT(w http.ResponseWriter, r *http.Request) {
	value := r.FormValue("someValue")
	io.WriteString(w, "Search " + value)
	// address/?someValue=searchPhrase
}
