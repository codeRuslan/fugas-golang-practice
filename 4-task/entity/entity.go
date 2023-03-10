package entity

type Book struct {
	Name   string `json:"name"`
	Author string `json:"author"`
	Year   Year   `json:"year"`
}

type SortedBooks []Book
