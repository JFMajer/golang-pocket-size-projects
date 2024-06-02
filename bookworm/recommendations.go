package main

type Recommendations struct {
	Name             string
	RecommendedBooks map[Book]uint
}

func GetRecommendations(bookworms []Bookworm) []Recommendations {
	// create new Recommendations slice
	var allRecommendations []Recommendations
	for i, bookworm := range bookworms {
		// add bookworm entry to Recommendations slice with emmpty map
		newReco := Recommendations{Name: bookworm.Name, RecommendedBooks: make(map[Book]uint)}
		for j, other_bookworm := range bookworms {
			if i == j {
				continue
			}
			// find common books to exlude them from being recommended
			commonBooksSlice := findCommonBooks([]Bookworm{bookworm, other_bookworm})
			commonBooks := make(map[Book]struct{})
			for _, book := range commonBooksSlice {
				commonBooks[book] = struct{}{}
			}
			// iterate over books on other_bookworm shelf, add +1 to recommended score
			for _, book := range other_bookworm.Books {
				// check if book is in commonBooks and exlude it from being added to recommendations
				if _, exists := commonBooks[book]; !exists {
					newReco.RecommendedBooks[book]++
				}
			}
		}
		allRecommendations = append(allRecommendations, newReco)
	}
	return allRecommendations
}
