package main

import (
	"net/http"
	"fmt"
	"log"
)

func main() {
	http.HandleFunc("/", set)
	http.HandleFunc("/read", read)
	http.HandleFunc("/second", second)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func set(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name: "userName",
		Value: "value",
	})
	fmt.Fprintln(w, "cookie set - ")
	fmt.Fprintln(w, "use dev tools")
}

func read(w http.ResponseWriter, r *http.Request) {
	c1, err := r.Cookie("userName")
	if err != nil {
		log.Println(err)
	} else {
		fmt.Fprintln(w, "Your cookie - ", c1)
	}

	c2, err := r.Cookie("general")
	if err != nil {
		log.Println(err)
	} else {
		fmt.Fprintln(w, "Your cookie - ", c2)
	}

	c3, err := r.Cookie("spec")
	if err != nil {
		log.Println(err)
	} else {
		fmt.Fprintln(w, "Your cookie - ", c3)
	}
}

func second (w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name: "general",
		Value: "some other general",
	})

	http.SetCookie(w, &http.Cookie{
		Name: "spec",
		Value: "value spec",
	})

	fmt.Fprintln(w, "Cookie set again")
}
