package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Seller struct {
	Name     string
	Image    string
	Phone    string
	Email    string
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

func loadSeller(title string) *Seller {
	filename := title + ".txt"
	s, _ := ioutil.ReadFile(filename)
	sellerInfo := strings.Split(string(s), ",")
	return &Seller{
		Name:     sellerInfo[0],
		Image:    sellerInfo[1],
		Phone:    sellerInfo[2],
		Location: sellerInfo[3],
	}
}

func (s *Seller) save() error {
	filename := s.Name + ".txt"
	return ioutil.WriteFile(filename, []byte(strings.Join([]string{s.Name, s.Image, s.Phone, s.Location}, ",")), 0600)
}
