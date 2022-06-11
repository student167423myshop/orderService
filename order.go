package main

import (
	"encoding/json"
	"errors"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

type Order struct {
	OrderId   string     `json:"OrderId"`
	ClientId  string     `json:"ClientId"`
	CartItems []CartItem `json:"CartItems"`
	Address   Address    `json:"Address"`
	TotalPaid Price      `json:"TotalPaid"`
}

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
		return Address{}, errors.New("got empty variable")
	}
	var address = Address{email, streetAddress, zipCode, city}
	return address, nil
}

func getOrderFromRedis(orderId string) Order {
	client := getRedis()
	var order Order
	jsonData, err := client.Get(orderId).Result()
	if err != nil {
		order = Order{}
	} else {
		_ = json.Unmarshal([]byte(jsonData), &order)
	}
	return order
}

func getWriteOrder(order Order) []byte {
	jsonData, _ := json.Marshal(order)
	return jsonData
}

func getNewOrder(
	clientId string,
	cartItems []CartItem,
	address Address,
	price Price,
) Order {
	orderId := uuid.NewV4().String()
	return Order{orderId, clientId, cartItems, address, price}
}
