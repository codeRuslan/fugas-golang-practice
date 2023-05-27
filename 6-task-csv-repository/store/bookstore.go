package store

import (
	"awesomeProject1/entity"
	"github.com/gocarina/gocsv"
	"os"
)

type Book interface {
	GetAll() ([]entity.Book, error)
	Update(books []entity.Book) ([]entity.Book, error)
}

type book struct {
	Books       []entity.Book
	CSVFilePath string
}

func NewBook(books []entity.Book, csvFilePath string) Book {
	return &book{Books: books,
		CSVFilePath: csvFilePath}
}

func (bl *book) Update(books []entity.Book) ([]entity.Book, error) {
	bl.Books = books
	file, err := os.OpenFile(bl.CSVFilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return nil, err
	}

	defer file.Close()
	if err := gocsv.MarshalWithoutHeaders(&books, file); err != nil {
		panic(err)
	}

	return bl.Books, nil
}

func (bl *book) GetAll() ([]entity.Book, error) {
	bookCSVfile, err := os.OpenFile(bl.CSVFilePath, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer bookCSVfile.Close()

	BooksCSV := []entity.Book{}

	if err := gocsv.UnmarshalFile(bookCSVfile, &BooksCSV); err != nil {
		panic(err)
	}

	return BooksCSV, nil
}
