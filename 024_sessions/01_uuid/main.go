package main

import (
	"net/http"
	"github.com/satori/go.uuid"
	"fmt"
)

func main() {
	http.HandleFunc("/", blankPage)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func blankPage(w http.ResponseWriter, r *http.Request)  {
	c, err := r.Cookie("session")
	if err != nil {
		id, _ := uuid.NewV4()
		c = &http.Cookie{
			Name: "session",
			Value: id.String(),
			//Secure: true,
			HttpOnly: true,
			Path: "/",
		}
		http.SetCookie(w, c)

	}
	fmt.Println(c)
}
