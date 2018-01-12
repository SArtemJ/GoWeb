package main

import (
	"net/http"
	"fmt"
)

type myType int

func (m myType) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("first-key", "from first Key")
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	fmt.Fprintln(w, "<h1> simple example</h1>")
}

func main() {
	var a myType
	http.ListenAndServe(":8080", a)
}
