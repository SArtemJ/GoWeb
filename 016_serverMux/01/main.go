package main

import (
	"net/http"
	"io"
)

type hotdog int

//handler interface
func (m hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/dog":
		io.WriteString(w, "dog value")
	case "/cat":
		io.WriteString(w, "cat value")
	}
}

func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}
