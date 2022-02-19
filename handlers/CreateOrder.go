package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/jamesthomasw/assignment_2/models"
)

func (h handler) CreateOrder(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	
	if err != nil {
		log.Fatalln(err)
	}

	var order models.Order
	json.Unmarshal(body, &order)

	if result := h.DB.Create(&order); result.Error != nil {
		fmt.Println(result.Error)
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Created")
}