package main

import (
	"net/http"
	"io"
	"os"
)

func example(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html; charset=UTF-8")

	io.WriteString(w, `
	<img src="test.jpg">
	`)
}

func serveMyFile(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open("test.jpg")
	if err != nil {
		http.Error(w, "file not fount", 404)
		return
	}
	defer file.Close()

	io.Copy(w, file)
}

func main() {
	http.HandleFunc("/", example)
	http.HandleFunc("/test.jpg", serveMyFile)
	http.ListenAndServe(":8080", nil)
}
