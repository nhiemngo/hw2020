package main

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/kr/pretty"
)

type sellerDB struct {
	DB *sql.DB
}

type Seller struct {
	Name     string
	Logo     string
	Image    string
	Phone    string
	Location string // can be used for map. or DaySchedule.Address can be used instead if we want to do multiple locations
	Schedule []*DaySchedule
}

type DaySchedule struct {
	Date      string
	Address   string // where the seller will be on this day
	StartTime string
	EndTime   string
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
	var logo string
	var image string
	var phone string
	var location string

	n_id, err := strconv.Atoi(id)
	if err != nil {
		return nil
	}

	address := "1455 Market St, San Francisco, CA 94103, USA"
	lat, lng := getGeocoding(address)
	pretty.Printf("Coordinates of location: %.7f, %.7f", lat, lng)

	row := sDB.DB.QueryRow(`SELECT name, logo, image, phone, location FROM seller WHERE id=?`, n_id)
	error := row.Scan(&name, &logo, &image, &phone, &location)

	if error != nil {
		return nil
	}
	return &Seller{
		Name:     name,
		Logo:     logo,
		Image:    image,
		Phone:    phone,
		Location: location,
	}
}

func (sDB *sellerDB) save(s Seller) int64 {
	res, err := sDB.DB.Exec(`INSERT INTO seller (
		name,
		logo,
		image,
		phone,
		location)
	VALUES (
		?,
		?,
		?,
		?,
		?
	);`, s.Name, s.Logo, s.Image, s.Phone, s.Location)
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

	return -1
}
