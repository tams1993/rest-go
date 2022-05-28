package main

import (
	"log"
	"net/http"

	"github.com/GolangProject/rest/pkg/db"
	"github.com/GolangProject/rest/pkg/handlers"
	"github.com/gorilla/mux"
)



func main() {
	DB := db.InitDB()
	h := handlers.New(DB)
	router := mux.NewRouter()
	router.HandleFunc("/books", h.GetAllBoooks).Methods(http.MethodGet)
	router.HandleFunc("/books/{id}", h.GetBookByID).Methods(http.MethodGet)
	router.HandleFunc("/books", h.AddBook).Methods(http.MethodPost)
	router.HandleFunc("/books/{id}", h.UpdateBookByID).Methods(http.MethodPut)
	router.HandleFunc("/books/{id}", h.DeleteBookByID).Methods(http.MethodDelete)


	log.Println("API is running on port 8080 !!!")
	log.Println(http.ListenAndServe(":8080", router))
}
