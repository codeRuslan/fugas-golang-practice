package store

import "awesomeProject1/entity"

type Book interface {
	GetAll() []entity.Book
	Update(books []entity.Book) ([]entity.Book, error)
}

type book struct {
	Books []entity.Book
}

func NewBook(books []entity.Book) Book {
	return &book{Books: books}
}

func (bl *book) GetAll() []entity.Book {
	return bl.Books
}

func (bl *book) Update(books []entity.Book) ([]entity.Book, error) {
	bl.Books = books
	return bl.Books, nil
}
