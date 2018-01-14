package main

import (
	"net/http"
	"io"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/test",  test)
	http.ListenAndServe(":8080", nil)
}

func test(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charest=utf-8)")
	io.WriteString(w, `<ing src="test.jpg"`)
}
