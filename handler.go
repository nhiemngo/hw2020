package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/skip2/go-qrcode"
)

const baseURL = "24cdf7b47d2d.ngrok.io"

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

func (app *application) viewHandler(w http.ResponseWriter, r *http.Request) {
	queryString := r.URL.Query()
	id := queryString.Get("id")
	s := app.db.loadSeller(id)
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

	sellerName := r.FormValue("Name")
	link := fmt.Sprintf("https://%v/view/?name=%v", baseURL, sellerName)
	qrcode.WriteFile(link, qrcode.Medium, 256, sellerName+"_qr.png")

	currentSeller := Seller{
		Name:     r.FormValue("Name"),
		Image:    r.FormValue("Image"),
		Phone:    r.FormValue("Phone"),
		Location: r.FormValue("Location"),
	}

	app.db.save(currentSeller)

	http.Redirect(w, r, "/view/?name="+r.FormValue("Name"), http.StatusFound)
}
