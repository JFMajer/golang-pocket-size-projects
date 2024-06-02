package main

import (
	"reflect"
	"sort"
	"testing"
)

type testCase struct {
	bookWormFile string
	want         []Bookworm
	wantErr      bool
}

var (
	handmaidsTale = Book{Author: "Margaret Atwood", Title: "The Handmaid's Tale"}
	oryxAndCrake  = Book{Author: "Margaret Atwood", Title: "Oryx and Crake"}
	theBellJar    = Book{Author: "Sylvia Plath", Title: "The Bell Jar"}
	janeEyre      = Book{Author: "Charlotte Brontë", Title: "Jane Eyre"}
)

func TestLoadBookworms(t *testing.T) {
	tests := map[string]testCase{
		"file exists": {
			bookWormFile: "testdata/bookworms.json",
			want: []Bookworm{
				{Name: "Fadi", Books: []Book{handmaidsTale, theBellJar}},
				{Name: "Peggy", Books: []Book{oryxAndCrake, handmaidsTale, janeEyre}},
			},
			wantErr: false,
		},
		"file does not exist": {
			bookWormFile: "testdata/bookworms2.json",
			want:         nil,
			wantErr:      true,
		},
		"invalid json": {
			bookWormFile: "testdata/bookworms3.json",
			want:         nil,
			wantErr:      true,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := loadBookworms(tc.bookWormFile)
			if err != nil && !tc.wantErr {
				t.Errorf("expected no error, got one %s", err.Error())
			}

			if err == nil && tc.wantErr {
				t.Errorf("expected error, got none")
			}

			if !equalBookworms(got, tc.want) {
				t.Errorf("different result: got %v, expected %v", got, tc.want)
			}
		})
	}

}

func equalBookworms(b1, b2 []Bookworm) bool {
	return reflect.DeepEqual(b1, b2)
}

// equalBooksCount is a helper to test the equality of two maps of books count.
func equalBooksCount(got, want map[Book]uint) bool {
	if len(got) != len(want) {
		return false
	}

	for book, targetCount := range want {
		count, ok := got[book]
		if !ok || targetCount != count {
			return false
		}
	}

	return true
}

func TestBooksCount(t *testing.T) {
	tt := map[string]struct {
		input []Bookworm
		want  map[Book]uint
	}{
		"nominal use case": {
			input: []Bookworm{
				{Name: "Fadi", Books: []Book{handmaidsTale, theBellJar}},
				{Name: "Peggy", Books: []Book{oryxAndCrake, handmaidsTale, janeEyre}},
			},
			want: map[Book]uint{handmaidsTale: 2, theBellJar: 1, oryxAndCrake: 1, janeEyre: 1},
		},
		"no bookworms": {
			input: []Bookworm{},
			want:  map[Book]uint{},
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got := booksCount(tc.input)
			if !equalBooksCount(tc.want, got) {
				t.Fatalf("got a different list of books: %v, expected %v", got, tc.want)
			}
		})
	}
}

func TestFindCommonBooks(t *testing.T) {
	tt := map[string]struct {
		input []Bookworm
		want  []Book
	}{
		"no common book": {
			input: []Bookworm{
				{Name: "Fadi", Books: []Book{handmaidsTale, theBellJar}},
				{Name: "Peggy", Books: []Book{oryxAndCrake, janeEyre}},
			},
			want: nil,
		},
		"one common book": {
			input: []Bookworm{
				{Name: "Fadi", Books: []Book{janeEyre, theBellJar}},
				{Name: "Peggy", Books: []Book{oryxAndCrake, janeEyre}},
			},
			want: []Book{janeEyre},
		},
		"three bookworms have the same books on their shelves": {
			input: []Bookworm{
				{Name: "Fadi", Books: []Book{janeEyre, theBellJar}},
				{Name: "Peggy", Books: []Book{theBellJar, janeEyre}},
				{Name: "Jakub", Books: []Book{theBellJar, janeEyre}},
			},
			want: []Book{janeEyre, theBellJar},
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got := findCommonBooks(tc.input)
			sortBooks(got)
			if !equalBooks(t, tc.want, got) {
				t.Fatalf("got a different list of books: %v, expected %v", got, tc.want)
			}
		})
	}

}

// equalBooks is a helper to test the equality of two lists of Books.
func equalBooks(t *testing.T, books, target []Book) bool {
	t.Helper()

	if len(books) != len(target) {
		// Early exit!
		return false
	}
	// Verify the content of the collections of Books for each Bookworm.
	for i := range target {
		if target[i] != books[i] {
			return false
		}
	}
	// Everything is equal!
	return true
}

func sortBooks(books []Book) {
	sort.Slice(books, func(i, j int) bool {
		if books[i].Author != books[j].Author {
			return books[i].Author < books[j].Author
		}
		return books[i].Title < books[j].Title
	})
}
