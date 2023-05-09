package repository

import (
	"awesomeProject1/entity"
	"github.com/gocarina/gocsv"
	"os"
)

//func main() {
//Test := GetCSVBooks("input_test_data.csv")
/*for _, book := range Test {
	fmt.Printf("Name: %s, Author: %s, Year: %d\n", book.Name, book.Author, book.Year)
}*/
//books := GetCSVBooks("/Users/ruslanpilipyuk/GolandProjects/awesomeProject1/repository/input_test_data.csv")
/*books := []entity.Book{
	entity.Book{Name: "TESTRAGE", Author: "Stephen King", Year: 1977},
	entity.Book{Name: "Philosopher's Stone", Author: "J. K. Rowling", Year: 1997},
	entity.Book{Name: "All Quiet on the Western Front", Author: "Erich Maria Remarque", Year: 1929},
}

test, _ := UpdateCSVBooks("/Users/ruslanpilipyuk/GolandProjects/awesomeProject1/repository/input_test_data.csv", books)
fmt.Println(test)*/
//}

func GetCSVBooks(path string) []entity.Book {
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

func UpdateCSVBooks(path string, books []entity.Book) ([]entity.Book, error) {
	file, _ := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)

	defer file.Close()
	if err := gocsv.MarshalWithoutHeaders(&books, file); err != nil {
		panic(err)
	}

	return books, nil
}
