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
	Name     string
	Image    string
	Phone    string
	Location string // dynamic location not yet supported???
	Schedule []*DaySchedule
}

type DaySchedule struct {
	Date      string
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
	var image string
	var phone string
	var location string

	n_id, err := strconv.Atoi(id)
	if err != nil {
		return nil
	}

	row := sDB.DB.QueryRow(`SELECT name, image, phone, location FROM seller WHERE id=?`, n_id)
	error := row.Scan(&name, &image, &phone, &location)
	if error != nil {
		return nil
	}
	return &Seller{
		Name:     name,
		Image:    image,
		Phone:    phone,
		Location: location,
	}
}

func (sDB *sellerDB)  save(s Seller) error {
	insert, err := sDB.DB.Query(`INSERT INTO seller (
		name,
		image,
		phone,
		location)
	VALUES (
		?,
		?,
		?,
		?
	);`, s.Name, s.Image, s.Phone, s.Location)
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
	return err
}
