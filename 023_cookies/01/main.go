package main

import (
	"net/http"
	"fmt"
)

func main() {
	http.HandleFunc("/", set)
	http.HandleFunc("/read", read)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func set(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name: "userName",
		Value: "value",
	})
	fmt.Fprintln(w, "Cookie written - check browser")
	fmt.Fprintln(w, "use dev tools")
}

func read(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("userName")
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	fmt.Fprintln(w, "Your cookie", c)
}