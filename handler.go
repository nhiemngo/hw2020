package main

import (
	"html/template"
	"net/http"
	"strings"
)

func functionHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("insert landing page"))
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	queryString := r.URL.Query()
	name := queryString.Get("name")
	s := loadSeller(name)
	w.Write([]byte(strings.Join([]string{s.name, s.img, s.phone, s.location}, ",")))
}

func createHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("/Users/nhiem/Development/go/src/sellercom/form.html"))
	if r.Method != http.MethodPost {
		tmpl.Execute(w, nil)
		return
	}
	currentSeller := Seller{
		name:     r.FormValue("name"),
		img:      r.FormValue("img"),
		phone:    r.FormValue("phone"),
		location: r.FormValue("location"),
	}

	currentSeller.save()

	http.Redirect(w, r, "/view/?name="+r.FormValue("name"), http.StatusFound)
	//tmpl.Execute(w, struct{ Success bool }{true})
}
