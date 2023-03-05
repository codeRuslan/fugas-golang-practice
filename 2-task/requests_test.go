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
	t.Run("Sucess ReturnAllBooks", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockBookStore := mock.NewMockBookStore(ctrl)
		mockBookStore.EXPECT().GetAllBooks().Return(ExpectedBooksGet)

		handler := handlers.Handler{BookStore: mockBookStore}

		req, err := http.NewRequest(http.MethodGet, "/books", nil)
		assert.NoError(t, err)

		rr := httptest.NewRecorder()

		handler.ReturnAllBooks(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)

		var books []entity.Book

		err = json.Unmarshal(rr.Body.Bytes(), &books)

		assert.NoError(t, err)

		assert.Equal(t, ExpectedBooksGet, books)

	})

	t.Run("Sucess CreateNewBooks", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockBookStore := mock.NewMockBookStore(ctrl)
		mockBookStore.EXPECT().CreateNewBooks(InputBooksPut).Return(ExpectedBooksPut, nil)

		handler := handlers.Handler{BookStore: mockBookStore}

		jsonInput, err := json.Marshal(InputBooksPut)
		assert.NoError(t, err)

		req, err := http.NewRequest(http.MethodPut, "/books", bytes.NewReader(jsonInput))
		assert.NoError(t, err)

		rr := httptest.NewRecorder()

		handler.CreateNewBook(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)

		var books []entity.Book

		err = json.Unmarshal(rr.Body.Bytes(), &books)
		assert.NoError(t, err)
		assert.Equal(t, ExpectedBooksPut, books)

	})
}
