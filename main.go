package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func getRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/healthz", healthzHandler).Methods(http.MethodGet)
	r.HandleFunc("/order/{userId}", getOrderHandler).Methods(http.MethodGet)
	r.HandleFunc("/order/{userId}", newOrderHandler).Methods(http.MethodPost)
	return r
}

func main() {
	r := getRouter()

	srv := &http.Server{
		Handler: r,
		Addr:    ":7071",
	}

	if err := srv.ListenAndServe(); err != nil {
		panic(err.Error())
	}
}
