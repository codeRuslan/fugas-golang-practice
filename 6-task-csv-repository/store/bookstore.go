package store

import (
	"awesomeProject1/entity"
	"github.com/gocarina/gocsv"
	"os"
)

type Book interface {
	GetAll() ([]entity.Book, error)
	Update(books []entity.Book) error
}

type book struct {
	CSVFilePath string
}

func NewBook(csvFilePath string) Book {
	return &book{
		CSVFilePath: csvFilePath,
	}
}

func (b *book) Update(books []entity.Book) error {
	file, err := os.OpenFile(b.CSVFilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return nil
	}

	defer file.Close()
	if err := gocsv.MarshalWithoutHeaders(&books, file); err != nil {
		panic(err)
	}

	return nil
}

func (b *book) GetAll() ([]entity.Book, error) {
	bookCSVfile, err := os.OpenFile(b.CSVFilePath, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer bookCSVfile.Close()

	books := []entity.Book{}

	if err := gocsv.UnmarshalFile(bookCSVfile, &books); err != nil {
		panic(err)
	}

	return books, nil
}
