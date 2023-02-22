package main

/*
To test PUT Method
curl -X PUT \
-H "Content-Type: application/json" \
-d  '{"name":"Rage", "author": "Stephen King", "year": "1000"}' \
localhost:10000/books/put
*/

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
	"strconv"
)

func main() {
	books := &BookList{
		Books: []Book{
			Book{Name: "Rage", Author: "Stephen King", Year: "1977"},
			Book{Name: "Philosopher's Stone", Author: "J. K. Rowling", Year: "1997"},
			Book{Name: "All Quiet on the Western Front", Author: "Erich Maria Remarque", Year: "1929"},
		},
	}
	sort.Sort(sortedBooks(books.Books))
	handleRequests(books)
}

func handleRequests(books BookStore) {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/books", func(w http.ResponseWriter, r *http.Request) {
		returnAllBooks(w, r, books)
	})
	myRouter.HandleFunc("/books/put", func(w http.ResponseWriter, r *http.Request) {
		CreateNewBook(w, r, books)
	}).Methods("PUT")

	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func returnAllBooks(w http.ResponseWriter, r *http.Request, books BookStore) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(books.GetAllBooks())
}

func CreateNewBook(w http.ResponseWriter, r *http.Request, books BookStore) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var book Book
	json.Unmarshal(reqBody, &book)
	books.AddBook(book)

	json.NewEncoder(w).Encode(book)
}

type Book struct {
	Name   string `json:"name"`
	Author string `json:"author"`
	Year   string `json:"year"`
}

type sortedBooks []Book

func (a sortedBooks) Len() int {
	return len(a)
}

func (a sortedBooks) Less(i, j int) bool {
	iInt, _ := strconv.Atoi(a[i].Year)
	jInt, _ := strconv.Atoi(a[j].Year)
	return iInt < jInt
}

func (a sortedBooks) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

type BookStore interface {
	GetAllBooks() []Book
	AddBook(book Book)
}

type BookList struct {
	Books []Book
}

func (bl *BookList) GetAllBooks() []Book {
	return bl.Books
}

func (bl *BookList) AddBook(book Book) {
	bl.Books = append(bl.Books, book)
}
