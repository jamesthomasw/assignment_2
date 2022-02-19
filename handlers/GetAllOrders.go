package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jamesthomasw/assignment_2/models"
)

func (h handler) GetAllOrders(w http.ResponseWriter, r *http.Request) {
	var order []models.Order

	if result := h.DB.Find(&order); result.Error != nil {
		fmt.Println(result.Error)
	}
	
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(order)
}