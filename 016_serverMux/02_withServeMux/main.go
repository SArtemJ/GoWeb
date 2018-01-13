package main

import (
	"net/http"
	"io"
)

type hotdog int
type hotcat int

//handler interface
func (d hotdog) ServeHTTP(res http.ResponseWriter, r *http.Request) {
		io.WriteString(res, "dog value")

}

func (c hotcat) ServeHTTP(res http.ResponseWriter, r *http.Request) {
	io.WriteString(res, "cat value")

}

func main() {
	var d hotdog
	var c hotcat

	mux := http.NewServeMux()
	mux.Handle("/dog", d)
	mux.Handle("/cat", c)
	http.ListenAndServe(":8080", mux)

	// use default ServerMux
	//http.Handle("/dog", d)
	//http.Handle("/cat", c)
	//or
	//func d (res http.ResponseWriter, r *http.Request) {
	//io.WriteString(res, "dog value")
	//
	//}
	//func c (res http.ResponseWriter, r *http.Request) {
	//io.WriteString(res, "dog value")
	//
	//}
	//http.HandleFunc("/dog", d)
	//http.HandleFunc("/cat", c)
	//
	//http.ListenAndServe(":8080", nil)
}
