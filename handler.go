package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"github.com/skip2/go-qrcode"
)

const baseURL = "24cdf7b47d2d.ngrok.io"

func functionHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("insert landing page"))
}

func viewTestSellerHandler(w http.ResponseWriter, r *http.Request) {
	s := testSeller()

	tmpl, err := template.ParseFiles("templates/display.html")
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

func (app *application) viewHandler(w http.ResponseWriter, r *http.Request) {
	queryString := r.URL.Query()
	id := queryString.Get("id")
	s := app.db.loadSeller(id)
	tmpl, err := template.ParseFiles("templates/display.html")
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

func (app *application) createHandler(w http.ResponseWriter, r *http.Request) {
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
		Logo:	  r.FormValue("Logo"),
		Image:    r.FormValue("Image"),
		Phone:    r.FormValue("Phone"),
		Location: r.FormValue("Location"),
	}

	id := app.db.save(currentSeller)
	id_str := strconv.Itoa(int(id))

	link := fmt.Sprintf("https://%v/view/?id=%v", baseURL, id_str)
	qrcode.WriteFile(link, qrcode.Medium, 256, id_str+"_qr.png")

	http.Redirect(w, r, "/view/?id="+id_str, http.StatusFound)
}
