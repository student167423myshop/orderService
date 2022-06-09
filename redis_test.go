package main

import (
	"testing"
)

func TestGetRedis(t *testing.T) {
	initMockRedis()
	defer redisMockServer.Close()
	client := getRedis()
	_, err := client.Ping().Result()
	if err != nil {
		t.Fatal(err)
	}
}
