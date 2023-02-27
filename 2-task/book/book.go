package book

type Year int

type Book struct {
	Name   string `json:"name"`
	Author string `json:"author"`
	Year   Year   `json:"year"`
}

type SortedBooks []Book

func (a SortedBooks) Len() int {
	return len(a)
}

func (a SortedBooks) Less(i, j int) bool {
	return a[i].Year < a[j].Year
}

func (a SortedBooks) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
