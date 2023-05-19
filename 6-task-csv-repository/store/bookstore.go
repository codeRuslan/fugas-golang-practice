package store

import (
	"awesomeProject1/config"
	"awesomeProject1/entity"
	"github.com/gocarina/gocsv"
	"os"
)

var ConfigFileBookstore *config.Config

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
	file, err := os.OpenFile(ConfigFileBookstore.FilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return nil, err
	}

	defer file.Close()
	if err := gocsv.MarshalWithoutHeaders(&books, file); err != nil {
		panic(err)
	}

	return bl.Books, nil
}

func ReadBooksInCSV(path string) []entity.Book {
	bookCSVfile, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer bookCSVfile.Close()

	BooksCSV := []entity.Book{}

	if err := gocsv.UnmarshalFile(bookCSVfile, &BooksCSV); err != nil {
		panic(err)
	}

	return BooksCSV
}