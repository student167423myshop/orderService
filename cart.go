package main

import (
	"os"

	"github.com/go-resty/resty/v2"
)

type Cart struct {
	ClientId  string     `json:"ClientId"`
	CartItems []CartItem `json:"CartItems"`
}

type CartItem struct {
	ProductId string `json:"ProductId"`
	Quantity  int    `json:"Quantity"`
}

func getCart(userId string) (Cart, error) {
	client := resty.New()
	addr := os.Getenv("CART_SERVICE_ADDR")
	if addr == "" {
		addr = "http://localhost:7070"
	}
	var cart Cart
	_, err := client.
		R().
		SetResult(&cart).
		Get(addr + "/cart/" + userId)
	return cart, err
}

func emptyCart(userId string) error {
	client := resty.New()
	addr := os.Getenv("CART_SERVICE_ADDR")
	if addr == "" {
		addr = "http://localhost:7070"
	}
	_, err := client.
		R().
		Get(addr + "/cart/" + userId + "/empty")
	return err
}
