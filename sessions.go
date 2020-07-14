package main

import (
	"fmt"
	"github.com/gorilla/sessions"
	"html/template"
	"net/http"
)

var (
	store       = sessions.NewCookieStore([]byte("secret key???"))
	sessionName = "current-session"
)

func login(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, sessionName)
	if err != nil {
		session, err = store.New(r, sessionName)
	}

	tmpl, err := template.ParseFiles("login.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if r.Method != http.MethodPost {
		tmpl.Execute(w, nil)
		return
	}

	// authenticate here
	fmt.Println(r.FormValue("Name"))
	fmt.Println(r.FormValue("Password"))

	// set user as authenticated
	session.Values["authenticated"] = true
	session.Save(r, w)
}

func logout(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, sessionName)
	if err != nil {
		fmt.Println(err)
		return
	}
	// Revoke users authentication
	session.Values["authenticated"] = false
	session.Save(r, w)
}

func secret(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, sessionName)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Session not initiated", http.StatusForbidden)
		return
	}

	// Check if user is authenticated
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	// Print secret message
	fmt.Fprintln(w, "The cake is a lie!")
}
