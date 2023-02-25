package bookstore

import "awesomeProject1/model"

type BookStore interface {
	GetAllBooks() []model.Book
	AddBook(book model.Book)
}

type BookList struct {
	Books []model.Book
}

func (bl *BookList) GetAllBooks() []model.Book {
	return bl.Books
}

func (bl *BookList) AddBook(book model.Book) {
	bl.Books = append(bl.Books, book)
}
