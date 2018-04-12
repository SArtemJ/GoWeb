package main

import (
	"github.com/satori/go.uuid"
	"net/http"
)

func getUser(w http.ResponseWriter, request *http.Request) user {
	c, err := request.Cookie("session")
	if err != nil {
		sID, _ := uuid.NewV4()
		c = &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
	}

	http.SetCookie(w, c)

	var u user
	if un, ok := dbSessions[c.Value]; ok {
		u = dbUsers[un]
	}
	return u
}

func alreadyLoggedIn(request *http.Request) bool {
	c, err := request.Cookie("session")
	if err != nil {
		return false
	}
	un := dbSessions[c.Value]
	_, ok := dbUsers[un]
	return ok
}
