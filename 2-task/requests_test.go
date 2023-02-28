package main

import (
	"awesomeProject1/book"
	"awesomeProject1/bookstore"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
)

func Test_get_method(t *testing.T) {

	books := &bookstore.BookList{
		Books: []book.Book{
			book.Book{Name: "Rage", Author: "Stephen King", Year: 1977},
			book.Book{Name: "Philosopher's Stone", Author: "J. K. Rowling", Year: 1997},
			book.Book{Name: "All Quiet on the Western Front", Author: "Erich Maria Remarque", Year: 1929},
		},
	}

	t.Run("Check GET Method", func(t *testing.T) {

		handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ReturnAllBooks(w, r, books)
		})

		req, err := http.NewRequest(http.MethodGet, "/books", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}

		/*expected := `[{"name":"All Quiet on the Western Front","author":"Erich Maria Remarque","year":1929},{"name":"Rage","author":"Stephen King","year":1977},{"name":"Philosopher's Stone","author":"J. K. Rowling","year":1997}]`
		if !reflect.DeepEqual(rr.Body.String(), expected) {
			t.Errorf("handler returned unexpected body: got %v want %v",
				rr.Body.String(), expected)
		}*/
	})

	t.Run("Check PUT Method", func(t *testing.T) {
		sample_data := `[{"name":"The Great Gatsby","author":"F. Scott Fitzgerald","year":1925}]`
		req, err := http.NewRequest("PUT", "/books", strings.NewReader(sample_data))
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Set("Content-Type", "application/json")
		recorder := httptest.NewRecorder()
		books := &bookstore.BookList{}
		CreateNewBook(recorder, req, books)

		if status := recorder.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}

		expectedBooks := []book.Book{
			book.Book{Name: "The Great Gatsby", Author: "F. Scott Fitzgerald", Year: 1925},
		}

		if !reflect.DeepEqual(books.Books, expectedBooks) {
			t.Errorf("handler returned unexpected body: got %v want %v",
				books.Books, expectedBooks)
		}

	})

}
