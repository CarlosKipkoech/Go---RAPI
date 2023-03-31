package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

//Book struct (Model)

type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

// author struct

type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var books []Book

//get All Books
func getBooks(responseWriter http.ResponseWriter, request *http.Request) {

}

func getBook(responseWriter http.ResponseWriter, request *http.Request) {

}
func createBook(responseWriter http.ResponseWriter, request *http.Request) {

}

func updateBook(responseWriter http.ResponseWriter, request *http.Request) {

}

func deleteBook(responseWriter http.ResponseWriter, request *http.Request) {

}

func main() {

	//@todo - implement DB (USING STATIC (HARD CODED))
	books = append(books, Book{ID: "1", Isbn: "349859", Title: "Book one", Author: &Author{
		Firstname: "John", Lastname: "Doe"}})

	books = append(books, Book{ID: "1", Isbn: "349859", Title: "Book one", Author: &Author{
		Firstname: "John", Lastname: "Doe"}})
	books = append(books, Book{ID: "1", Isbn: "349859", Title: "Book one", Author: &Author{
		Firstname: "John", Lastname: "Doe"}})
	books = append(books, Book{ID: "1", Isbn: "349859", Title: "Book one", Author: &Author{
		Firstname: "John", Lastname: "Doe"}})
	books = append(books, Book{ID: "1", Isbn: "349859", Title: "Book one", Author: &Author{
		Firstname: "John", Lastname: "Doe"}})
	router := mux.NewRouter()
	router.HandleFunc("/api/books", getBooks).Methods("GET")
	router.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/api/books", createBook).Methods("POST")
	router.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	router.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

}
