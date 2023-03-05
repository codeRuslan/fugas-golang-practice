package main

import "awesomeProject1/entity"

var (
	ExpectedBooksGet = []entity.Book{
		{Name: "Rage", Author: "Stephen King", Year: entity.Year(1977)},
		{Name: "Philosopher's Stone", Author: "J. K. Rowling", Year: entity.Year(1977)},
		{Name: "All Quiet on the Western Front", Author: "Erich Maria Remarque", Year: entity.Year(1929)},
	}

	ExpectedBooksPut = []entity.Book{
		{Name: "Rage", Author: "Stephen King", Year: entity.Year(1977)},
		{Name: "Philosopher's Stone", Author: "J. K. Rowling", Year: entity.Year(1977)},
		{Name: "All Quiet on the Western Front", Author: "Erich Maria Remarque", Year: entity.Year(1929)},
		{Name: "The Fellowship of the Ring", Author: "J. R. R. Tolkien", Year: entity.Year(1954)},
	}

	InputBooksPut = []entity.Book{
		{Name: "The Fellowship of the Ring", Author: "J. R. R. Tolkien", Year: entity.Year(1954)},
	}
)
