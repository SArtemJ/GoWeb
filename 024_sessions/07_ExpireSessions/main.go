package main

import (
	"github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	"html/template"
	"net/http"
	"time"
)

type user struct {
	UserName string
	Password []byte
	First    string
	Last     string
	Role string
}

type session struct {
	 un string
	 lastActivity time.Time
}

var tpl *template.Template
var dbUsers = map[string]user{}      // user ID, user
var dbSessions = map[string]session{} // session ID, session
var dbSessionsCleaned time.Time
const sessionLength int = 30

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
	//bs, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.MinCost)
	//dbUsers["test@test.com"] = user{"test@test.com", bs, "James", "Bond"}
	dbSessionsCleaned = time.Now()
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, request *http.Request) {
	u := getUser(w, request)
	showSessions()
	tpl.ExecuteTemplate(w, "index.gohtml", u)
}

func bar(w http.ResponseWriter, request *http.Request) {
	u := getUser(w, request)
	if !alreadyLoggedIn(w, request) {
		http.Redirect(w, request, "/", http.StatusSeeOther)
		return
	}
	if u.Role != "007" {
		http.Error(w, "You must be 007", http.StatusForbidden)
		return
	}
	tpl.ExecuteTemplate(w, "bar.gohtml", u)
}

func signup(w http.ResponseWriter, request *http.Request) {
	if alreadyLoggedIn(w, request) {
		http.Redirect(w, request, "/", http.StatusSeeOther)
		return
	}
	var u user
	// process form submission
	if request.Method == http.MethodPost {
		// get form values
		un := request.FormValue("username")
		p := request.FormValue("password")
		f := request.FormValue("firstname")
		l := request.FormValue("lastname")
		r := request.FormValue("role")
		// username taken?
		if _, ok := dbUsers[un]; ok {
			http.Error(w, "Username already taken", http.StatusForbidden)
			return
		}
		// create session
		sID, _ := uuid.NewV4()
		c := &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		c.MaxAge = sessionLength
		http.SetCookie(w, c)
		dbSessions[c.Value] = session{un, time.Now()}
		// store user in dbUsers
		bs, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.MinCost)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		u = user{un, bs, f, l, r}
		dbUsers[un] = u
		// redirect
		http.Redirect(w, request, "/", http.StatusSeeOther)
		return
	}
	showSessions()
	tpl.ExecuteTemplate(w, "signup.gohtml", u)
}

func login(w http.ResponseWriter, request *http.Request) {
	if alreadyLoggedIn(w, request) {
		http.Redirect(w, request, "/", http.StatusSeeOther)
		return
	}

	var u user
	// process form submission
	if request.Method == http.MethodPost {
		un := request.FormValue("username")
		p := request.FormValue("password")
		// is there a username?
		u, ok := dbUsers[un]
		if !ok {
			http.Error(w, "Username and/or password do not match", http.StatusForbidden)
			return
		}
		// does the entered password match the stored password?
		err := bcrypt.CompareHashAndPassword(u.Password, []byte(p))
		if err != nil {
			http.Error(w, "Username and/or password do not match", http.StatusForbidden)
			return
		}
		// create session
		sID, _ := uuid.NewV4()
		c := &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		c.MaxAge = sessionLength
		http.SetCookie(w, c)
		dbSessions[c.Value] = session{un, time.Now()}
		http.Redirect(w, request, "/", http.StatusSeeOther)
		return
	}
	showSessions()
	tpl.ExecuteTemplate(w, "login.gohtml", u)
}

func logout(w http.ResponseWriter, request *http.Request) {
	if !alreadyLoggedIn(w, request) {
		http.Redirect(w, request, "/", http.StatusSeeOther)
		return
	}
	c, _ := request.Cookie("session")
	delete(dbSessions, c.Value)

	c = &http.Cookie{
		Name:   "session",
		Value:  "",
		MaxAge: -1,
	}
	http.SetCookie(w, c)

	if time.Now().Sub(dbSessionsCleaned) > (time.Second * 30 ){
		go cleanSessions()
	}

	http.Redirect(w, request, "/login", http.StatusSeeOther)

}
