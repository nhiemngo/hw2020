package main

import (
	"io/ioutil"
	"strings"
)

type Seller struct {
	name     string
	img      string
	phone    string
	location string
}

func loadSeller(title string) *Seller {
	filename := title + ".txt"
	s, _ := ioutil.ReadFile(filename)
	sellerInfo := strings.Split(string(s), ",")
	return &Seller{
		name:     sellerInfo[0],
		img:      sellerInfo[1],
		phone:    sellerInfo[2],
		location: sellerInfo[3],
	}
}

func (s *Seller) save() error {
	filename := s.name + ".txt"
	return ioutil.WriteFile(filename, []byte(strings.Join([]string{s.name, s.img, s.phone, s.location}, ",")), 0600)
}

type allSellers struct {
	sellers []Seller
}
