package  main

import (
	"net/http"
	"fmt"
)


func main()  {

http.HandleFunc("/", blankPage)
http.HandleFunc("/star", star)
http.Handle("/favicon.ico", http.NotFoundHandler())
http.ListenAndServe(":8080", nil)

}

func blankPage(w http.ResponseWriter, r *http.Request) {
fmt.Print("Request method blankPage is ", r.Method, "\n\n")
}

func star(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request method star is ", r.Method)
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}
