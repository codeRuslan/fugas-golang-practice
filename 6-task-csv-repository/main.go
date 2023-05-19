package main

import (
	"awesomeProject1/config"
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

	store.ConfigFileBookstore = ConfigFile
	books := store.ReadBooksInCSV(ConfigFile.FilePath)
	bookstore := store.NewBook(books)
	handlers.HandleRequests(bookstore)
}
