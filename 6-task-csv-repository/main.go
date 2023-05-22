package main

import (
	"awesomeProject1/config"
	"awesomeProject1/entity"
	"awesomeProject1/handlers"
	"awesomeProject1/store"
	"log"
)

var ConfigFile config.Config

func main() {
	ConfigFile, err := config.ReadJsonConfigFile("/Users/ruslanpilipyuk/GolandProjects/awesomeProject1/config/config.json")
	if err != nil {
		log.Fatal(err)
	}

	bookstore := store.NewBook([]entity.Book{}, ConfigFile.FilePath)

	books, err := bookstore.ReadBooksInCSV()
	if err != nil {
		panic(err)
	}
	
	bookstore = store.NewBook(books, ConfigFile.FilePath)
	handlers.HandleRequests(bookstore)
}
