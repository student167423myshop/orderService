package main

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/mux"
)

func healthzHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func getOrderHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orderId := vars["orderId"]
	order := getOrderFromRedis(orderId)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(getWriteOrder(order))
}

func newOrderHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	clientId := vars["userId"]
	address, err := getAddressFromForm(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	cart, err := getCart(clientId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	order, err := addOrderToRedis(cart, address)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		emptyCart(clientId)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(getWriteOrder(order))
	}
}

func addOrderToRedis(cart Cart, address Address) (Order, error) {
	shippingCost := getShippingCost()
	cartItems := cart.CartItems
	productIds := getProductIds(cartItems)
	products, err := getProducts(productIds)
	if err != nil {
		panic(err.Error())
	}
	productsPrice := getProductsPrice(products, cartItems)
	price := getTotalPrice(shippingCost, productsPrice)
	order := getNewOrder(cart.ClientId, cart.CartItems, address, price)
	json, err := json.Marshal(order)
	if err != nil {
		return Order{}, errors.New(err.Error())
	}
	client := getRedis()
	err = client.Set(order.OrderId, json, 0).Err()
	if err != nil {
		return Order{}, errors.New(err.Error())
	}
	return getOrderFromRedis(order.OrderId), nil
}
