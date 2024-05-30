package main

import (
	"encoding/json"
	"os"
)

type Bookworm struct {
	Name  string `json:"name"`
	Books []Book `json:"books"`
}

type Book struct {
	Author string `json:"author"`
	Title  string `json:"title"`
}

func loadBookworms(filepath string) ([]Bookworm, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var bookworms []Bookworm
	err = json.NewDecoder(file).Decode(&bookworms)
	if err != nil {
		return nil, err
	}

	return bookworms, nil
}

func findCommonBooks(bookworms []Bookworm) []Book {
	books := make(map[Book]int)
	for _, bkwms := range bookworms {
		for _, book := range bkwms.Books {
			books[book]++

		}
	}
	var booksToReturn []Book
	for k, v := range books {
		if v > 1 {
			booksToReturn = append(booksToReturn, k)
		}
	}
	return booksToReturn
}
