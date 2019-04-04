package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"restapi/book"
)

//books
var books []book.Book

func getBooks(c echo.Context) error {
	return c.JSON(http.StatusOK, books)
}

func getBook(c echo.Context) error {
	id := c.Param("id")

	for _, item := range books {
		if item.ID == id {
			return c.JSON(http.StatusOK, item)
		}
	}

	return c.JSON(http.StatusNotFound,"Not Found")
}
//
//func createBook(c echo.Context) error) {
//	w.Header().Set("Content-type", "application/json")
//	var book book.Book
//	_ = json.NewDecoder(r.Body).Decode(&book)
//	book.ID = strconv.Itoa(rand.Intn(10000000)) //not safe obviously
//	books = append(books, book)
//
//	json.NewEncoder(w).Encode(book)
//}

//func updateBook(w http.ResponseWriter, r *http.Request) {
//
//}
//
//func deleteBook(w http.ResponseWriter, r *http.Request) {
//
//}

func main() {
	//init router

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	books = append(books, book.Book{ID: "1", Isbn: "543543", Title: "Book One", Author: &book.Author{Firstname: "John", Lastname: "Doe"}})
	books = append(books, book.Book{ID: "2", Isbn: "12345", Title: "Book Two", Author: &book.Author{Firstname: "Jack", Lastname: "Don't"}})

	e.GET("/api/books", getBooks)
	e.GET("/api/books/:id", getBook)
	//e.POST("/api/books", createBook)
	//router.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	//router.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	e.Logger.Fatal(e.Start(":8000"))
}
