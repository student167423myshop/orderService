package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_Healthz(t *testing.T) {
	r := getRouter()
	mockServer := httptest.NewServer(r)
	resp, _ := http.Get(mockServer.URL + "/healthz")
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Status should be ok, got %d", resp.StatusCode)
	}
}
