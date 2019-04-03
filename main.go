package main

import (
	"encoding/json"
	"log"
	"net/http"
	"math/rand"
	"strconv"

	"github.com/gorilla/mux"
)

func main() {
	//init router

	router := mux.NewRouter()

	router.HandleFunc("/api/books", getBooks).Methods("GET")
	router.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/api/books",createBook).Methods("POST")
	router.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	router.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")
}
