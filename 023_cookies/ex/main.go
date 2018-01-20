package main

import (
	"net/http"
	"fmt"
	"log"
	"strconv"
)

var counter int

func main() {
	http.HandleFunc("/", set)
	http.HandleFunc("/times", howManyTimes)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func set(w http.ResponseWriter, r *http.Request){
	counter += 1
	http.SetCookie(w, &http.Cookie{
		Name: "userName",
		Value: strconv.Itoa(counter),
	} )
}

func howManyTimes(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("userName")

	//set default cookie
	//if err == http.ErrNoCookie{
	//	c = &http.Cookie{
	//		Name: "userName",
	//		Value: "0",
	//	}
	//}

	//strconv.Atoi - ascii to int

	if err != nil {
		log.Println(err)
	} else {
		fmt.Fprintln(w, "How many times - ", c)
	}
}
