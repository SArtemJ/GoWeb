package main

import (
	"net/http"
	"github.com/satori/go.uuid"
	"log"
)

func getUser(w http.ResponseWriter, req *http.Request) user{

	c, err := req.Cookie("session")
	log.Printf("cookie c get user %v", c)

	if err != nil {
		sID, _ := uuid.NewV4()
		c = &http.Cookie{
			Name: "session",
			Value: sID.String(),

		}
	}

	var u user
	log.Printf("User get user %v", u)
	if un, ok := dbSessions[c.Value]; ok {
		u = dbUsers[un]
	}
	log.Printf("User get user %v", u)
	return  u
}


func alreadyLoggedIn(req *http.Request) bool {

	c, err := req.Cookie("session")
	if err != nil {
		return  false
	}
	un := dbSessions[c.Value]
	_, ok := dbUsers[un]
	return ok

}

