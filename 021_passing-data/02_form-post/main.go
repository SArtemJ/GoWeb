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
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	io.WriteString(w, `<form method="post">
	<input type="text" name="someValue">
	<input type="submit">
	</form>
	<br>` + value)
}
