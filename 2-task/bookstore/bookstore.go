package bookstore

import "awesomeProject1/entity"

type BookStore interface {
	GetAllBooks() []entity.Book
	CreateNewBooks(books []entity.Book) ([]entity.Book, error)
}

type bookStore struct {
	Books []entity.Book
}

func (bl *bookStore) GetAllBooks() []entity.Book {
	return bl.Books
}

func NewBookList(books []entity.Book) BookStore {
	return &bookStore{Books: books}
}

func (bl *bookStore) CreateNewBooks(books []entity.Book) ([]entity.Book, error) {
	bl.Books = books
	return bl.Books, nil
}
