package main

import (
	"encoding/json"

	uuid "github.com/satori/go.uuid"
)

type Order struct {
	OrderId   string     `json:"OrderId"`
	ClientId  string     `json:"ClientId"`
	CartItems []CartItem `json:"CartItems"`
	Address   Address    `json:"Address"`
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

func getNewOrder(clientId string, cartItems []CartItem, address Address) Order {
	orderId := uuid.NewV4().String()
	return Order{orderId, clientId, cartItems, address}
}
