package main

import (
	"database/sql"
	"fmt"
	"strconv"
)

type sellerDB struct {
	DB *sql.DB
}

type Seller struct {
	Name      string
	About     string
	Logo      string
	Image     string
	Image_2   string
	Image_3   string
	Image_4   string
	Image_5   string
	Phone     string
	Location  string // can be used for map. or DaySchedule.Address can be used instead if we want to do multiple locations
	Email     string
	Twitter   string
	Facebook  string
	Instagram string
	Pinterest string
	Schedule  []*DaySchedule
}

type DaySchedule struct {
	Date       string
	OpenStatus string
	Location   string
	Address    string // where the seller will be on this day
	StartTime  string
	EndTime    string
}

func testSeller() *Seller {
	var ds []*DaySchedule

	mondaySched := &DaySchedule{
		Date:      "Monday",
		StartTime: "5am",
		EndTime:   "10pm",
	}

	tuesdaySched := &DaySchedule{
		Date:      "Tuesday",
		StartTime: "6am",
		EndTime:   "9pm",
	}

	ds = append(ds, mondaySched)
	ds = append(ds, tuesdaySched)

	s := &Seller{
		Name:     "testName",
		Image:    "testImage",
		Phone:    "123456",
		Location: "testAddress",
		Schedule: ds,
	}

	fmt.Println(s)
	fmt.Println(s.Schedule[0])

	return s
}

func (sDB *sellerDB) loadSeller(id string) *Seller {
	var name string
	var about string
	var logo string
	var image string
	var image_2 string
	var image_3 string
	var image_4 string
	var image_5 string
	var phone string
	var location string
	var email string
	var twitter string
	var facebook string
	var instagram string
	var pinterest string
	var ds []*DaySchedule

	var mon_status bool
	var mon_location string
	var mon_address string
	var mon_start string
	var mon_end string

	var tue_status bool
	var tue_location string
	var tue_address string
	var tue_start string
	var tue_end string

	var wed_status bool
	var wed_location string
	var wed_address string
	var wed_start string
	var wed_end string

	var thu_status bool
	var thu_location string
	var thu_address string
	var thu_start string
	var thu_end string

	var fri_status bool
	var fri_location string
	var fri_address string
	var fri_start string
	var fri_end string

	var sat_status bool
	var sat_location string
	var sat_address string
	var sat_start string
	var sat_end string

	var sun_status bool
	var sun_location string
	var sun_address string
	var sun_start string
	var sun_end string

	n_id, err := strconv.Atoi(id)
	if err != nil {
		return nil
	}

	row := sDB.DB.QueryRow(`SELECT monday, 
									monday_location, 
									monday_address, 
									monday_start, 
									monday_end FROM schedule WHERE id=?`, n_id)
	error := row.Scan(&mon_status, &mon_location, &mon_address, &mon_start, &mon_end)
	if error != nil {
		return nil
	}

	row = sDB.DB.QueryRow(`SELECT tuesday, 
									tuesday_location, 
									tuesday_address, 
									tuesday_start, 
									tuesday_end FROM schedule WHERE id=?`, n_id)
	error = row.Scan(&tue_status, &tue_location, &tue_address, &tue_start, &tue_end)
	if error != nil {
		return nil
	}

	row = sDB.DB.QueryRow(`SELECT wednesday, 
									wednesday_location, 
									wednesday_address, 
									wednesday_start, 
									wednesday_end FROM schedule WHERE id=?`, n_id)
	error = row.Scan(&wed_status, &wed_location, &wed_address, &wed_start, &wed_end)
	if error != nil {
		return nil
	}

	row = sDB.DB.QueryRow(`SELECT thursday, 
									thursday_location, 
									thursday_address, 
									thursday_start, 
									thursday_end FROM schedule WHERE id=?`, n_id)
	error = row.Scan(&thu_status, &thu_location, &thu_address, &thu_start, &thu_end)
	if error != nil {
		return nil
	}

	row = sDB.DB.QueryRow(`SELECT friday, 
									friday_location, 
									friday_address, 
									friday_start, 
									friday_end FROM schedule WHERE id=?`, n_id)
	error = row.Scan(&fri_status, &fri_location, &fri_address, &fri_start, &fri_end)
	if error != nil {
		return nil
	}

	row = sDB.DB.QueryRow(`SELECT saturday, 
									saturday_location, 
									saturday_address, 
									saturday_start, 
									saturday_end FROM schedule WHERE id=?`, n_id)
	error = row.Scan(&sat_status, &sat_location, &sat_address, &sat_start, &sat_end)
	if error != nil {
		return nil
	}

	row = sDB.DB.QueryRow(`SELECT sunday, 
									sunday_location, 
									sunday_address, 
									sunday_start, 
									sunday_end FROM schedule WHERE id=?`, n_id)
	error = row.Scan(&sun_status, &sun_location, &sun_address, &sun_start, &sun_end)
	if error != nil {
		return nil
	}

	monSched := &DaySchedule{
		Date:      "mon",
		Location:  mon_location,
		Address:   mon_address,
		StartTime: mon_start,
		EndTime:   mon_end,
	}

	tueSched := &DaySchedule{
		Date:      "tue",
		Location:  tue_location,
		Address:   tue_address,
		StartTime: tue_start,
		EndTime:   tue_end,
	}

	wedSched := &DaySchedule{
		Date:      "wed",
		Location:  wed_location,
		Address:   wed_address,
		StartTime: wed_start,
		EndTime:   wed_end,
	}

	thuSched := &DaySchedule{
		Date:      "thu",
		Location:  thu_location,
		Address:   thu_address,
		StartTime: thu_start,
		EndTime:   thu_end,
	}

	friSched := &DaySchedule{
		Date:      "fri",
		Location:  fri_location,
		Address:   fri_address,
		StartTime: fri_start,
		EndTime:   fri_end,
	}
	satSched := &DaySchedule{
		Date:      "sat",
		Location:  sat_location,
		Address:   sat_address,
		StartTime: sat_start,
		EndTime:   sat_end,
	}
	sunSched := &DaySchedule{
		Date:      "sun",
		Location:  sun_location,
		Address:   sun_address,
		StartTime: sun_start,
		EndTime:   sun_end,
	}

	ds = append(ds, monSched)
	ds = append(ds, tueSched)
	ds = append(ds, wedSched)
	ds = append(ds, thuSched)
	ds = append(ds, friSched)
	ds = append(ds, satSched)
	ds = append(ds, sunSched)

	row = sDB.DB.QueryRow(`SELECT name, 
									about,
									logo, 
									image, 
									second_image, 
									third_image,
									fourth_image,
									fifth_image,
									phone, 
									location,
									email,
									twitter,
									facebook,
									instagram,
									pinterest FROM seller WHERE id=?`, n_id)
	error = row.Scan(&name, &about, &logo, &image, &image_2, &image_3, &image_4, &image_5, &phone, &location,
		&email, &twitter, &facebook, &instagram, &pinterest)
	if error != nil {
		return nil
	}
	return &Seller{
		Name:      name,
		About:     about,
		Logo:      logo,
		Image:     image,
		Image_2:   image_2,
		Image_3:   image_3,
		Image_4:   image_4,
		Image_5:   image_5,
		Phone:     phone,
		Location:  location,
		Email:     email,
		Twitter:   twitter,
		Facebook:  facebook,
		Instagram: instagram,
		Pinterest: pinterest,
		Schedule:  ds,
	}
}

func (sDB *sellerDB) save(s Seller) int64 {
	insert, err := sDB.DB.Query(`INSERT INTO seller (
		name,
		about
		logo,
		image,
		second_image,
		third_image,
		fourth_image,
		fifth_image,
		phone,
		location,
		email,
		twitter,
		facebook,
		instagram,
		pinterest
		)
	VALUES (
		?,
		?,
		?,
		?,
		?,
		?,
		?,
		?,
		?,
		?,
		?,
		?,
		?,
		?,
		?
	);`, s.Name, s.About, s.Logo, s.Image, s.Image_2, s.Image_3, s.Image_4, s.Image_5,
		s.Phone, s.Location, s.Email, s.Twitter, s.Facebook, s.Instagram, s.Pinterest)
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()

	monbool, _ := strconv.ParseBool(s.Schedule[0].OpenStatus)
	tuebool, _ := strconv.ParseBool(s.Schedule[1].OpenStatus)
	wedbool, _ := strconv.ParseBool(s.Schedule[2].OpenStatus)
	thubool, _ := strconv.ParseBool(s.Schedule[3].OpenStatus)
	fribool, _ := strconv.ParseBool(s.Schedule[4].OpenStatus)
	satbool, _ := strconv.ParseBool(s.Schedule[5].OpenStatus)
	sunbool, _ := strconv.ParseBool(s.Schedule[6].OpenStatus)

	res, err := sDB.DB.Exec(`INSERT INTO schedule (
		monday,
		tuesday,
		wednesday,
		thursday,
		friday,
		saturday,
		sunday,
		monday_start,
		tuesday_start,
		wednesday_start,
		thursday_start,
		friday_start,
		saturday_start,
		sunday_start,
		monday_end,
		tuesday_end,
		wednesday_end,
		thursday_end,
		friday_end,
		saturday_end,
		sunday_end,
		monday_location,
		tuesday_location,
		wednesday_location,
		thursday_location,
		friday_location,
		saturday_location,
		sunday_location,
		monday_address,
		tuesday_address,
		wednesday_address,
		thursday_address,
		friday_address,
		saturday_address,
		sunday_address)
	VALUES (
		?,
		?,
		?,
		?,
		?,
		?,
		?,
		?,
		?,
		?,
		?,
		?,
		?,
		?,
		?,
		?,
		?,
		?,
		?,
		?,
		?,
		?,
		?,
		?,
		?,
		?,
		?,
		?,
		?,
		?,
		?,
		?,
		?,
		?,
		?
	);`, monbool, tuebool, wedbool, thubool, fribool, satbool, sunbool, s.Schedule[0].StartTime, s.Schedule[1].StartTime, s.Schedule[2].StartTime, s.Schedule[3].StartTime,
		s.Schedule[4].StartTime, s.Schedule[5].StartTime, s.Schedule[6].StartTime, s.Schedule[0].EndTime, s.Schedule[1].EndTime, s.Schedule[2].EndTime,
		s.Schedule[3].EndTime, s.Schedule[4].EndTime, s.Schedule[5].EndTime, s.Schedule[6].EndTime, s.Schedule[0].Location, s.Schedule[1].Location,
		s.Schedule[2].Location, s.Schedule[3].Location, s.Schedule[4].Location, s.Schedule[5].Location, s.Schedule[6].Location, s.Schedule[0].Address, s.Schedule[1].Address,
		s.Schedule[2].Address, s.Schedule[3].Address, s.Schedule[4].Address, s.Schedule[5].Address, s.Schedule[6].Address)
	if err != nil {
		println("Exec err:", err.Error())
	} else {
		id, err := res.LastInsertId()
		if err != nil {
			println("Error:", err.Error())
		} else {
			return id
		}
	}

	return 0
}
