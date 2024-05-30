package main

import (
	"fmt"
	"log"
)

func main() {
	bookworms, err := loadBookworms("./testdata/bookworms.json")
	if err != nil {
		log.Fatal("Something went wrong: ", err)
	}

	for _, b := range bookworms {
		fmt.Printf("Name: %s\n", b.Name)
		for _, book := range b.Books {
			fmt.Printf("  Book: %s by %s\n", book.Title, book.Author)
		}
	}

	commonBooks := findCommonBooks(bookworms)
	fmt.Println("Common books:")
	for _, book := range commonBooks {
		fmt.Printf("Author: %s, Title: %s\n", book.Author, book.Title)
	}
}
