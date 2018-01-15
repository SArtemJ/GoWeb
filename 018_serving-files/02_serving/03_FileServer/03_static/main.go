package main

import (
	"net/http"
)

func main() {
	//if not index.html, access to directory
	http.ListenAndServe(":8080", http.FileServer(http.Dir(".")))
}
