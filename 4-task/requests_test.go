package main

import (
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
	t.Run("Sucess Get Existing Books", func(t *testing.T) {
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

		assert.Equal(t, "{\"books\":[{\"name\":\"Rage\",\"author\":\"Stephen King\",\"year\":1977},{\"name\":\"Philosopher's Stone\",\"author\":\"J. K. Rowling\",\"year\":1977},{\"name\":\"All Quiet on the Western Front\",\"author\":\"Erich Maria Remarque\",\"year\":1929}],\"date\":\"05.03.2023\"}\n", rr.Body.String())

	})

	t.Run("Sucess Put New Books", func(t *testing.T) {
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

		assert.Equal(t, "{\"books\":[{\"name\":\"Rage\",\"author\":\"Stephen King\",\"year\":1977},{\"name\":\"Philosopher's Stone\",\"author\":\"J. K. Rowling\",\"year\":1977},{\"name\":\"All Quiet on the Western Front\",\"author\":\"Erich Maria Remarque\",\"year\":1929},{\"name\":\"The Fellowship of the Ring\",\"author\":\"J. R. R. Tolkien\",\"year\":1954}],\"date\":\"05.03.2023\"}\n", rr.Body.String())

	})
}
