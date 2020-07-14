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
	Location string
}

func loadSeller(title string) *Seller {
	filename := title + ".txt"
	s, _ := ioutil.ReadFile(filename)
	sellerInfo := strings.Split(string(s), ",")
	fmt.Println("*****")
	fmt.Println(sellerInfo)
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

type allSellers struct {
	sellers []Seller
}
