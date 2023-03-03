package main

import (
	"awesomeProject1/entity"
	"awesomeProject1/handlers"
	"awesomeProject1/mock"
	"bytes"
	"encoding/json"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestBooksGet(t *testing.T) {
	t.Run("Check GET Method", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expectedBooks := []entity.Book{
			{Name: "Rage", Author: "Stephen King", Year: entity.Year(1977)},
			{Name: "Philosopher's Stone", Author: "J. K. Rowling", Year: entity.Year(1977)},
			{Name: "All Quiet on the Western Front", Author: "Erich Maria Remarque", Year: entity.Year(1929)},
		}

		mockBookStore := mock.NewMockBookStore(ctrl)
		mockBookStore.EXPECT().GetAllBooks().Return(expectedBooks)

		handler := handlers.Handler{BookStore: mockBookStore}

		req, err := http.NewRequest(http.MethodGet, "/books", nil)
		assert.NoError(t, err)

		rr := httptest.NewRecorder()

		handler.ReturnAllBooks(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)

		var books []entity.Book

		err = json.Unmarshal(rr.Body.Bytes(), &books)

		assert.NoError(t, err)

		assert.Equal(t, expectedBooks, books)

	})

	t.Run("Check PUT Method", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expectedBooks := []entity.Book{
			{Name: "Rage", Author: "Stephen King", Year: entity.Year(1977)},
			{Name: "Philosopher's Stone", Author: "J. K. Rowling", Year: entity.Year(1977)},
			{Name: "All Quiet on the Western Front", Author: "Erich Maria Remarque", Year: entity.Year(1929)},
			{Name: "The Fellowship of the Ring", Author: "J. R. R. Tolkien", Year: entity.Year(1954)},
		}

		inputBooks := []entity.Book{
			{Name: "The Fellowship of the Ring", Author: "J. R. R. Tolkien", Year: entity.Year(1954)},
		}

		mockBookStore := mock.NewMockBookStore(ctrl)
		mockBookStore.EXPECT().CreateNewBooks(inputBooks).Return(expectedBooks, nil)

		handler := handlers.Handler{BookStore: mockBookStore}

		jsonInput, err := json.Marshal(inputBooks)
		assert.NoError(t, err)

		req, err := http.NewRequest(http.MethodPut, "/books", bytes.NewReader(jsonInput))
		assert.NoError(t, err)

		rr := httptest.NewRecorder()

		handler.CreateNewBook(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)

		var books []entity.Book

		err = json.Unmarshal(rr.Body.Bytes(), &books)
		assert.NoError(t, err)
		assert.Equal(t, expectedBooks, books)

	})
}
