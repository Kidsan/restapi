package main

import (
	"encoding/json"
	"log"
	"net/http"
//	"math/rand"
//	"strconv"
	book"restapi/book"

	"github.com/gorilla/mux"
)

//books

var books []book.Book

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type","application/json")
	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {

}

func createBook(w http.ResponseWriter, r *http.Request) {

}

func updateBook(w http.ResponseWriter, r *http.Request) {

}

func deleteBook(w http.ResponseWriter, r *http.Request) {

}

func main() {
	//init router

	router := mux.NewRouter()

	books = append(books, book.Book{ID: "1", Isbn: "543543", Title: "Book One", Author: &book.Author{ Firstname: "John", Lastname: "Doe"}})
	books = append(books, book.Book{ID: "2", Isbn: "12345", Title: "Book Two", Author: &book.Author{ Firstname: "Jack", Lastname: "Don't"}})

	router.HandleFunc("/api/books", getBooks).Methods("GET")
	router.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/api/books",createBook).Methods("POST")
	router.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	router.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}
