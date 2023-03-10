package main

import (
	"awesomeProject1/entity"
	"awesomeProject1/handlers"
	"awesomeProject1/store"
)

func main() {
	books := []entity.Book{
		entity.Book{Name: "Rage", Author: "Stephen King", Year: 1977},
		entity.Book{Name: "Philosopher's Stone", Author: "J. K. Rowling", Year: 1997},
		entity.Book{Name: "All Quiet on the Western Front", Author: "Erich Maria Remarque", Year: 1929},
	}

	bookstore := store.NewBook(books)
	handlers.HandleRequests(bookstore)
}
