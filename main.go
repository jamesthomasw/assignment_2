package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jamesthomasw/assignment_2/db"
	"github.com/jamesthomasw/assignment_2/handlers"
)

func main() {
	DB := db.Init()
	h := handlers.New(DB)
	r := mux.NewRouter()

	r.HandleFunc("/orders", h.GetAllOrders).Methods(http.MethodGet)
	r.HandleFunc("/orders/{id}", h.GetOrder).Methods(http.MethodGet)
	r.HandleFunc("/orders", h.CreateOrder).Methods(http.MethodPost)
	r.HandleFunc("/orders/{id}", h.UpdateOrder).Methods(http.MethodPut)
	r.HandleFunc("/orders/{id}", h.DeleteOrder).Methods(http.MethodDelete)
	
	log.Println("Connected to API!")
	http.ListenAndServe(":8080", r)
}