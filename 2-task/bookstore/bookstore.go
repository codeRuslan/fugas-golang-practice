package bookstore

import "awesomeProject1/book"

type BookStore interface {
	GetAllBooks() []book.Book
}

type bookList struct {
	Books []book.Book
}

func NewBookList(books []book.Book) BookStore {
	return &bookList{Books: books}
}

func (bl *bookList) GetAllBooks() []book.Book {
	return bl.Books
}
