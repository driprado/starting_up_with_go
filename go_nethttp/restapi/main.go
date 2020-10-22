// https://youtu.be/SonwZ6MF5BE?t=1174 ==> Do this from personal computer

package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
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

// Type ResponseWriter is an interface used by an HTTP handler to construct an HTTP response.
func getBooks(w http.ResponseWriter, r *http.Request) {

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
	//Initiare router
	router := mux.NewRouter()

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
