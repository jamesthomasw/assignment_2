package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jamesthomasw/assignment_2/models"
)

func (h handler) DeleteOrder(w http.ResponseWriter, r *http.Request) {
	
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	var order models.Order
	if result := h.DB.First(&order, id); result.Error != nil {
		fmt.Println(result.Error)
	}

	h.DB.Delete(&order)

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Deleted")
}