package handlers

import (
	"awesomeProject1/entity"
	"awesomeProject1/store"
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
)

type Handler struct {
	BookStore store.Book
}

func HandleRequests(books store.Book) {
	myRouter := mux.NewRouter().StrictSlash(true)
	handlerInstance := Handler{BookStore: books}
	myRouter.HandleFunc("/books", handlerInstance.ReturnAllBooks).Methods(http.MethodGet)
	myRouter.HandleFunc("/books", handlerInstance.CreateNewBook).Methods(http.MethodPut)

	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func (booksHandler *Handler) ReturnAllBooks(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	allBooks := booksHandler.BookStore.GetAll()
	json.NewEncoder(w).Encode(allBooks)
}

func (booksHandler *Handler) CreateNewBook(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}

	var newBooks []entity.Book
	if err := json.Unmarshal(reqBody, &newBooks); err != nil {
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}

	sort.Sort(entity.SortedBooks(newBooks))
	allBooks, err := booksHandler.BookStore.Update(newBooks)

	if err != nil {
		http.Error(w, "Failed to retrieve books", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(allBooks)
}
