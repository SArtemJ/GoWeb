package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"github.com/SArtemJ/GoWeb/MongoDB/03_PostDelete/models"
	"github.com/gin-gonic/gin/json"
	"fmt"
)

func main() {
	r := httprouter.New()
	r.GET("/", index)
	r.GET("/user/:id", getUser)
	r.POST("/user", createUser)
	r.DELETE("/user/:id", deleteUser)
	http.ListenAndServe("localhost:8080", r)
}

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	s := `
<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="utf-8">
<title>Index</title>
</head>
<body>
<a href="/user/987654321">GO TO: http://localhost:8080//user/987654321</a>
</body>
</html>
`


	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(s))
}

func getUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	u := models.User{
		Name:"TestUser",
		Gender:"male",
		Age: 28,
		Id: p.ByName("id"),
	}

	uj, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", uj)

}

func createUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	u := models.User{}

	json.NewDecoder(r.Body).Decode(&u)

	u.Id = "001"
	uj, _ := json.Marshal(u)

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s\n", uj)
}

func deleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Something")
}

