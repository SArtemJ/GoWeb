package main

import (
	"net/http"
	"io"
)

/*ListenAndServe on port ":8080" using the default ServeMux.

Use HandleFunc to add the following routes to the default ServeMux:

"/" "/dog/" "/me/

Add a func for each of the routes.

Have the "/me/" route print out your name. */

func indexPage(res http.ResponseWriter, r *http.Request) {
	io.WriteString(res, "Blank page")
}

func dogPage(res http.ResponseWriter, r *http.Request) {
	io.WriteString(res, "Dog page")
}

func mePage(res http.ResponseWriter, r *http.Request)  {
	io.WriteString(res, "Timon")
}

func main()  {

	//srvMux
	srvMux := http.NewServeMux()
	srvMux.HandleFunc("/", indexPage)
	srvMux.HandleFunc("/dog/", dogPage)
	srvMux.HandleFunc("/me/", mePage)
	http.ListenAndServe(":8080", srvMux)

	//default
	//http.HandleFunc("/", indexPage)
	//http.HandleFunc("/dog/", dogPage)
	//http.HandleFunc("/me/", mePage)
	//http.ListenAndServe(":8080", nil)

}
