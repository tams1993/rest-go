package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"github.com/GolangProject/rest/pkg/models"
	"github.com/gorilla/mux"
)

func (h handler) GetBookByID(w http.ResponseWriter, r *http.Request) {

	// Read dynamic ID parameter from Request
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"]) 


	// Find book by ID

	var book models.Book

	if result := h.DB.First(&book, id); result.Error != nil {
		fmt.Println(result.Error)
	}


	w.WriteHeader(http.StatusOK)
			w.Header().Add("Content-Type", "application/json")
			json.NewEncoder(w).Encode(book)

	 
}