package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/GolangProject/rest/pkg/models"
	"github.com/gorilla/mux"
)

func (h handler) UpdateBookByID(w http.ResponseWriter, r *http.Request) {
	// Read dynamic ID from parameter
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	// Read json request body
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var updatedBook models.Book
	json.Unmarshal(body, &updatedBook)

	var book models.Book

	if result := h.DB.First(&book, id); result.Error != nil {
		fmt.Println(result.Error)
	}

	book.Title = updatedBook.Title
	book.Author = updatedBook.Author
	book.Desc = updatedBook.Desc

	h.DB.Save(&book)

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
	// Update and send response when book Id matches dynamic Id
}
