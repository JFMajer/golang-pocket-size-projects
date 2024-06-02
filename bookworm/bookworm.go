package main

import (
	"bufio"
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

	bufferReader := bufio.NewReaderSize(file, 1024*1024)

	var bookworms []Bookworm
	err = json.NewDecoder(bufferReader).Decode(&bookworms)
	if err != nil {
		return nil, err
	}

	return bookworms, nil
}

func findCommonBooks(bookworms []Bookworm) []Book {
	books := booksCount(bookworms)
	var booksToReturn []Book
	for k, v := range books {
		if v > 1 {
			booksToReturn = append(booksToReturn, k)
		}
	}
	return booksToReturn
}

func booksCount(bookworms []Bookworm) map[Book]uint {
	books := make(map[Book]uint)
	for _, bookworm := range bookworms {
		for _, book := range bookworm.Books {
			books[book]++

		}
	}
	return books
}
