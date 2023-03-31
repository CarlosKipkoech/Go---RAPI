package main

import (
	"encoding/json"
	"log"
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

	//setContentType
	responseWriter.Header().Set("Content-Type", "application/json")
	json.NewEncoder(responseWriter).Encode(books)

}

func getBook(responseWriter http.ResponseWriter, request *http.Request) {

	responseWriter.Header().Set("Content-Type", "application/json")
	// get params
	params := mux.Vars(request)

	//loop thro the param to find with id
	for _, item := range books {

		if item.ID == params["id"] {

			json.NewEncoder(responseWriter).Encode(item)
			return

		}

	}

	// if no matching id is not found - encode empty Json

	json.NewEncoder(responseWriter).Encode(&Book{})

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

	books = append(books, Book{ID: "2", Isbn: "949857", Title: "Book Two", Author: &Author{
		Firstname: "Jane", Lastname: "Smith"}})
	books = append(books, Book{ID: "3", Isbn: "849856", Title: "Book Three", Author: &Author{
		Firstname: "Wesley", Lastname: "Vandem"}})
	books = append(books, Book{ID: "4", Isbn: "749856", Title: "Book four", Author: &Author{
		Firstname: "Swazz", Lastname: "Negear"}})
	books = append(books, Book{ID: "5", Isbn: "149855", Title: "Book Five", Author: &Author{
		Firstname: "Pusle", Lastname: "Roped"}})

	// initialize mux router handler
	router := mux.NewRouter()
	router.HandleFunc("/api/books", getBooks).Methods("GET")
	router.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/api/books", createBook).Methods("POST")
	router.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	router.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))

}
