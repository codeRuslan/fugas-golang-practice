package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
)

type Year int

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

//func (y int) MarshalJSON() ([]byte, error) {
//	return []byte(fmt.Sprintf(`"%d"`, y)), nil
//}

func main() {
	books := &BookList{
		Books: []Book{
			Book{Name: "Rage", Author: "Stephen King", Year: 1977},
			Book{Name: "Philosopher's Stone", Author: "J. K. Rowling", Year: 1997},
			Book{Name: "All Quiet on the Western Front", Author: "Erich Maria Remarque", Year: 1929},
		},
	}
	sort.Sort(sortedBooks(books.Books))
	handleRequests(books)
}

func handleRequests(books BookStore) {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/books", func(w http.ResponseWriter, r *http.Request) {
		ReturnAllBooks(w, r, books)
	}).Methods(http.MethodGet)
	myRouter.HandleFunc("/books/put", func(w http.ResponseWriter, r *http.Request) {
		CreateNewBook(w, r, books.(*BookList))
	}).Methods(http.MethodPut)

	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func ReturnAllBooks(w http.ResponseWriter, r *http.Request, books BookStore) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(books.GetAllBooks())
}

func CreateNewBook(w http.ResponseWriter, r *http.Request, books *BookList) {
	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}

	var newBooks []Book
	if err := json.Unmarshal(reqBody, &newBooks); err != nil {
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(newBooks)
}

type Book struct {
	Name   string `json:"name"`
	Author string `json:"author"`
	Year   Year   `json:"year"`
}

type sortedBooks []Book

func (a sortedBooks) Len() int {
	return len(a)
}

func (a sortedBooks) Less(i, j int) bool {
	//iInt, _ := strconv.Atoi(a[i].Year)
	//jInt, _ := strconv.Atoi(a[j].Year)
	return i < j
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
