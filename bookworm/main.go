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
	displayBooks(commonBooks)

	// Generate recommendations
	recommendations := GetRecommendations(bookworms)
	// Print recommendations
	for _, rec := range recommendations {
		fmt.Printf("Recommendations for %s:\n", rec.Name)
		for book, score := range rec.RecommendedBooks {
			fmt.Printf(" - %s by %s (Score: %d)\n", book.Title, book.Author, score)
		}
	}
}

func displayBooks(books []Book) {
	for _, b := range books {
		fmt.Println("-", b.Title, "by", b.Author)
	}
}
