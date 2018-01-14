package main

import (
	"io"
	"net/http"
	"os"
)

func example(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html; charset=UTF-8")

	io.WriteString(w, `
	<img src="/test.jpg">
	`)
}

func serveMyFile(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open("test.jpg")
	if err != nil {
		http.Error(w, "file not fount", 404)
		return
	}
	defer file.Close()

	fileE, err := file.Stat()
	if err != nil {
		http.Error(w, "file not found", 404)
		return
	}

	http.ServeContent(w, r, file.Name(), fileE.ModTime(), file)
}

func main() {
	http.HandleFunc("/", example)
	http.HandleFunc("/test.jpg", serveMyFile)
	http.ListenAndServe(":8080", nil)
}
