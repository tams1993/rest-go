package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"github.com/GolangProject/rest/pkg/models"
)

 

func (h handler) AddBook(w http.ResponseWriter, r *http.Request){

	// Read request body
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	// Convert to JSON
	var book models.Book
	json.Unmarshal(body, &book)

	// Append to the Book mocks
	if result := h.DB.Create(&book); result.Error != nil {

		fmt.Println(result.Error)

	}

	// Send a 201 created response
	w.WriteHeader(http.StatusCreated)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Created")

}