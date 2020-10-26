// https://youtu.be/SonwZ6MF5BE
// @TODO:
// 1 - Make delete from slice work - ongoing, reading slices docs (0_resources)
// 2 - Implement SQL DB into it
// 3 - Implement dynamoDB into it

package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

// Book Struct
type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

// Author Struct
type Author struct {
	Firstname string `json:"fistname"`
	Lastname  string `json:"lastname"`
}

// Initialise books variable as a slice of Book structures
// A slice is analogous to arrays, but have some unusual properties.
// https://blog.golang.org/slices-intro
var books []Book

// Get All books
func getBooks(w http.ResponseWriter, r *http.Request) { // Type ResponseWriter is an interface used by an HTTP handler to construct an HTTP response.
	w.Header().Set("Content-Type", "application/json") // Set ResponseWriter header to json
	json.NewEncoder(w).Encode(books)                   // ResponseWrier to encode our slice of books
}

// Get single book
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// Get parameters from URL
	parameters := mux.Vars(r) //func Vars(r *http.Request) map[string]string | returns the url variables for the current request.
	// Loop through books until find ID == idparam
	for _, book := range books {
		if book.ID == parameters["id"] {
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{}) // Encode everything in the Book structure

}

func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newbook Book
	_ = json.NewDecoder(r.Body).Decode(&newbook)  // Decode newbook request
	newbook.ID = strconv.Itoa(rand.Intn(1000000)) // Generate random ID number and convert to string
	books = append(books, newbook)
	json.NewEncoder(w).Encode(newbook) // Encode newbook response

}

// Update book
func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = params["id"]
			books = append(books, book)
			json.NewEncoder(w).Encode(book)
			return
		}
	}
}

// Delete book
func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	parameters := mux.Vars(r)
	for index, book := range books {
		if book.ID == parameters["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}

func main() {
	//Initiare router
	router := mux.NewRouter()

	// Mock Data --> @TODO: implement DB

	books = append(books, Book{ID: "1", Isbn: "438227", Title: "Book One", Author: &Author{Firstname: "John", Lastname: "Doe"}})
	books = append(books, Book{ID: "2", Isbn: "454555", Title: "Book Two", Author: &Author{Firstname: "Steve", Lastname: "Smith"}})

	// Route Handlers / Endpoints
	// func HandleFunc(pattern string, handler func(ResponseWriter, *Request))
	// HandleFunc registers the handler function for the given pattern in the DefaultServeMux.
	router.HandleFunc("/api/books", getBooks).Methods("GET")     // Get all books
	router.HandleFunc("/api/books/{id}", getBook).Methods("GET") // Get single book by id
	router.HandleFunc("/api/books", createBook).Methods("POST")
	router.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	router.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	// func ListenAndServe(addr string, handler Handler) error
	// apparently router is a handler as well
	log.Fatal(http.ListenAndServe(":8000", router)) // wrapped in log.Fatal to log in case of error

}
