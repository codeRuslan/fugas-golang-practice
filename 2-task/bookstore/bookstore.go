package bookstore

import "awesomeProject1/book"

type BookStore interface {
	GetAllBooks() []book.Book
}

type BookList struct {
	Books []book.Book
}

func (bl *BookList) GetAllBooks() []book.Book {
	return bl.Books
}
