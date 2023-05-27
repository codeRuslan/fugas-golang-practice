package handlers

import (
	"awesomeProject1/config"
	"awesomeProject1/entity"
	"awesomeProject1/store"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
	"time"
)

type Handler struct {
	BookStore store.Book
}

func HandleRequests(books store.Book) {
	myRouter := mux.NewRouter().StrictSlash(true)
	handlerInstance := Handler{BookStore: books}
	myRouter.HandleFunc("/books", handlerInstance.ReturnAllBooks).Methods(http.MethodGet)
	myRouter.HandleFunc("/books", handlerInstance.CreateNewBook).Methods(http.MethodPut)
	configFile, err := config.ReadJsonConfigFile("/Users/ruslanpilipyuk/GolandProjects/awesomeProject1/config/config.json")
	if err != nil {
		fmt.Println("Error when reading config file")
	}
	log.Fatal(http.ListenAndServe(configFile.ListenPort, myRouter))
}

func (booksHandler *Handler) ReturnAllBooks(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	allBooks, err := booksHandler.BookStore.GetAll()

	if err != nil {
		fmt.Println(err)
	}

	booksResp := entity.BookResponse{
		Books: allBooks,
		Date:  entity.CivilTime(time.Now()),
	}

	json.NewEncoder(w).Encode(booksResp)
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

	booksResp := entity.BookResponse{
		Books: allBooks,
		Date:  entity.CivilTime(time.Now()),
	}

	json.NewEncoder(w).Encode(booksResp)
}
