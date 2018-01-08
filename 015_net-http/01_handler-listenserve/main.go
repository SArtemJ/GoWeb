package main

import (
	"fmt"
	"net/http"
)

type hotdog int

//handler interface
func (m hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "My code")
}

func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}
