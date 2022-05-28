package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/GolangProject/rest/pkg/models"
	"github.com/gorilla/mux"
)

func (h handler) DeleteBookByID(w http.ResponseWriter, r *http.Request) {

	// Read the dynamic ID from parameter
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	// Find Book by ID

	var book models.Book

	if result := h.DB.First(&book, id); result.Error != nil {
		fmt.Println(result.Error)
	}

	h.DB.Delete(&book)

	// Delete Book 

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Deleted")

}
