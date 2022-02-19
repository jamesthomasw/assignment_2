package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jamesthomasw/assignment_2/models"
)

func (h handler) UpdateOrder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var updatedOrder models.Order
	json.Unmarshal(body, &updatedOrder)


	var order models.Order
	if result := h.DB.First(&order, id); result.Error != nil {
		fmt.Println(result.Error)
	}

	order.OrderedAt = updatedOrder.OrderedAt
	order.CustomerName = updatedOrder.CustomerName
	order.Items = updatedOrder.Items

	h.DB.Save(&order)

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Updated")
}