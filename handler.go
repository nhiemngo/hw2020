package main

import (
	"fmt"
	"github.com/skip2/go-qrcode"
	"html/template"
	"net/http"
	"strconv"
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

	var ds []*DaySchedule

	tmpl, err := template.ParseFiles("templates/form.html")

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if r.Method != http.MethodPost {
		tmpl.Execute(w, nil)
		return
	}

	fmt.Println("form submitted!")

	mondaySched := &DaySchedule{
		Date:       "Monday",
		OpenStatus: r.FormValue("MonOpenStatus"),
		Location:   r.FormValue("MonLocation"),
		Address:    r.FormValue("MonAddress"),
		StartTime:  r.FormValue("MonStartTime"),
		EndTime:    r.FormValue("MonEndTime"),
	}
	tuesdaySched := &DaySchedule{
		Date:       "Tuesday",
		OpenStatus: r.FormValue("TueOpenStatus"),
		Location:   r.FormValue("TueLocation"),
		Address:    r.FormValue("TueAddress"),
		StartTime:  r.FormValue("TueStartTime"),
		EndTime:    r.FormValue("TueEndTime"),
	}
	wednesdaySched := &DaySchedule{
		Date:       "Wednesday",
		OpenStatus: r.FormValue("WedOpenStatus"),
		Location:   r.FormValue("WedLocation"),
		Address:    r.FormValue("WedAddress"),
		StartTime:  r.FormValue("WedStartTime"),
		EndTime:    r.FormValue("WedEndTime"),
	}
	thursdaySched := &DaySchedule{
		Date:       "Thursday",
		OpenStatus: r.FormValue("ThuOpenStatus"),
		Location:   r.FormValue("ThuLocation"),
		Address:    r.FormValue("ThuAddress"),
		StartTime:  r.FormValue("ThuStartTime"),
		EndTime:    r.FormValue("ThuEndTime"),
	}
	fridaySched := &DaySchedule{
		Date:       "Friday",
		OpenStatus: r.FormValue("FriOpenStatus"),
		Location:   r.FormValue("FriLocation"),
		Address:    r.FormValue("FriAddress"),
		StartTime:  r.FormValue("FriStartTime"),
		EndTime:    r.FormValue("FriEndTime"),
	}
	saturdaySched := &DaySchedule{
		Date:       "Saturday",
		OpenStatus: r.FormValue("SatOpenStatus"),
		Location:   r.FormValue("SatLocation"),
		Address:    r.FormValue("SatAddress"),
		StartTime:  r.FormValue("SatStartTime"),
		EndTime:    r.FormValue("SatEndTime"),
	}
	sundaySched := &DaySchedule{
		Date:       "Sunday",
		OpenStatus: r.FormValue("SunOpenStatus"),
		Location:   r.FormValue("SunLocation"),
		Address:    r.FormValue("SunAddress"),
		StartTime:  r.FormValue("SunStartTime"),
		EndTime:    r.FormValue("SunEndTime"),
	}

	ds = append(ds, mondaySched)
	ds = append(ds, tuesdaySched)
	ds = append(ds, wednesdaySched)
	ds = append(ds, thursdaySched)
	ds = append(ds, fridaySched)
	ds = append(ds, saturdaySched)
	ds = append(ds, sundaySched)

	currentSeller := Seller{
		Name:      r.FormValue("Name"),
		About:	   r.FormValue("About"),
		Logo:      r.FormValue("Logo"),
		Image:     r.FormValue("Image"),
		Image_2:   r.FormValue("Image_2"),
		Image_3:   r.FormValue("Image_3"),
		Image_4:   r.FormValue("Image_4"),
		Image_5:   r.FormValue("Image_5"),
		Phone:     r.FormValue("Phone"),
		Location:  r.FormValue("Location"),
		Email:     r.FormValue("Email"),
		Twitter:   r.FormValue("Twitter"),
		Facebook:  r.FormValue("Facebook"),
		Instagram: r.FormValue("Instagram"),
		Pinterest: r.FormValue("Pinterest"),
		Schedule:  ds,
	}

	id := app.db.save(currentSeller)
	id_str := strconv.Itoa(int(id))

	link := fmt.Sprintf("https://%v/view/?id=%v", baseURL, id_str)
	qrcode.WriteFile(link, qrcode.Medium, 256, id_str+"_qr.png")

	http.Redirect(w, r, "/view/?id="+id_str, http.StatusFound)
}
