package main

import (
	"awesomeProject1/config"
	"awesomeProject1/handlers"
	"awesomeProject1/store"
)

func main() {
	configFile, err := config.ReadJsonConfigFile("/Users/ruslanpilipyuk/GolandProjects/awesomeProject1/config/config.json")
	if err != nil {
		panic(err)
	}
	books := store.ReadBooksInCSV(configFile.FilePath)
	bookstore := store.NewBook(books)
	handlers.HandleRequests(bookstore)
}
