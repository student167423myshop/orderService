package main

import (
	"errors"
	"net/http"
)

type Address struct {
	Email         string `json:"Email"`
	StreetAddress string `json:"StreetAddress"`
	ZipCode       string `json:"ZipCode"`
	City          string `json:"City"`
}

func getAddressFromForm(r *http.Request) (Address, error) {
	var email = r.FormValue("email")
	var streetAddress = r.FormValue("street_address")
	var zipCode = r.FormValue("zip_code")
	var city = r.FormValue("city")
	if email == "" || streetAddress == "" || zipCode == "" || city == "" {
		return Address{}, errors.New("Got empty variable")
	}
	var address = Address{email, streetAddress, zipCode, city}
	return address, nil
}
