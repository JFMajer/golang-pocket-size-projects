package main

import (
	"fmt"
	"log"
	"sort"
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
	printReccomendations(recommendations)
}

func displayBooks(books []Book) {
	for _, b := range books {
		fmt.Println("-", b.Title, "by", b.Author)
	}
}

func printReccomendations(recommendations []Recommendations) {
	for _, rec := range recommendations {
		fmt.Printf("Recommendations for %s:\n", rec.Name)

		sortedBooks := make([]Book, 0, len(rec.RecommendedBooks))
		for book := range rec.RecommendedBooks {
			sortedBooks = append(sortedBooks, book)
		}

		sort.Slice(sortedBooks, func(i, j int) bool {
			return rec.RecommendedBooks[sortedBooks[i]] > rec.RecommendedBooks[sortedBooks[j]]
		})

		for _, book := range sortedBooks {
			score := rec.RecommendedBooks[book]
			if score > 1 {
				fmt.Printf(" - %s by %s (Score: %d)\n", book.Title, book.Author, score)
			}
		}

	}
}
