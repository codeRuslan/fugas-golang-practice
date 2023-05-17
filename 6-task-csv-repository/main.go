package main

import (
	"awesomeProject1/config"
	"awesomeProject1/handlers"
	"awesomeProject1/repository"
	"awesomeProject1/store"
)

func main() {
	/*books := []entity.Book{
		entity.Book{Name: "Rage", Author: "Stephen King", Year: 1977}, Rage,Stephen King,1977
		entity.Book{Name: "Philosopher's Stone", Author: "J. K. Rowling", Year: 1997}, Philosopher's Stone,J. K. Rowling,1997
		entity.Book{Name: "All Quiet on the Western Front", Author: "Erich Maria Remarque", Year: 1929}, All Quiet on the Western Front,Erich Maria Remarque,1929
	}*/
	configFile, _ := config.ReadJsonConfigFile("/Users/ruslanpilipyuk/GolandProjects/awesomeProject1/config/config.json")
	books := repository.GetCSVBooks(configFile.FilePath)
	bookstore := store.NewBook(books)
	handlers.HandleRequests(bookstore)
}
