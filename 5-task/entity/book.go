package entity

import "encoding/json"

type Book struct {
	Name   string `json:"name"`
	Author string `json:"author"`
	Year   Year   `json:"year"`
}

type BookResponse struct {
	Books []Book    `json:"books"`
	Date  CivilTime `json:"date"`
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

func (y Year) MarshalJSON() ([]byte, error) {
	return json.Marshal(int(y))
}

func (y *Year) UnmarshalJSON(data []byte) error {
	var yearInt int
	if err := json.Unmarshal(data, &yearInt); err != nil {
		return err
	}
	*y = Year(yearInt)
	return nil
}
