package main

import (
	"awesomeProject1/bookstore"
	"awesomeProject1/model"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
)

type Year model.Year

func (y Year) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%d"`, y)), nil
}

func (y *Year) UnmarshalJSON(data []byte) error {
	var yearInt int
	if err := json.Unmarshal(data, &yearInt); err != nil {
		return err
	}
	*y = Year(yearInt)
	return nil
}

func main() {
	books := &bookstore.BookList{
		Books: []model.Book{
			model.Book{Name: "Rage", Author: "Stephen King", Year: 1977},
			model.Book{Name: "Philosopher's Stone", Author: "J. K. Rowling", Year: 1997},
			model.Book{Name: "All Quiet on the Western Front", Author: "Erich Maria Remarque", Year: 1929},
		},
	}
	sort.Sort(model.SortedBooks(books.Books))
	handleRequests(books)
}

func handleRequests(books bookstore.BookStore) {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/books", func(w http.ResponseWriter, r *http.Request) {
		ReturnAllBooks(w, r, books)
	}).Methods(http.MethodGet)
	myRouter.HandleFunc("/books/put", func(w http.ResponseWriter, r *http.Request) {
		CreateNewBook(w, r, books.(*bookstore.BookList))
	}).Methods(http.MethodPut)

	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func ReturnAllBooks(w http.ResponseWriter, r *http.Request, books bookstore.BookStore) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(books.GetAllBooks())
}

func CreateNewBook(w http.ResponseWriter, r *http.Request, books *bookstore.BookList) {
	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}

	var newBooks []model.Book
	if err := json.Unmarshal(reqBody, &newBooks); err != nil {
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(newBooks)
}
