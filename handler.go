package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func functionHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("insert landing page"))
}

func viewTestSellerHandler(w http.ResponseWriter, r *http.Request) {
	s := testSeller()

	tmpl, err := template.ParseFiles("display.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, s)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	queryString := r.URL.Query()
	name := queryString.Get("name")
	s := loadSeller(name)

	tmpl, err := template.ParseFiles("display.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, s)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func createHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("form.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if r.Method != http.MethodPost {
		tmpl.Execute(w, nil)
		return
	}

	fmt.Println("form submitted!")
	currentSeller := Seller{
		Name:     r.FormValue("Name"),
		Image:    r.FormValue("Image"),
		Phone:    r.FormValue("Phone"),
		Location: r.FormValue("Location"),
	}

	currentSeller.save()

	http.Redirect(w, r, "/view/?name="+r.FormValue("Name"), http.StatusFound)
}
