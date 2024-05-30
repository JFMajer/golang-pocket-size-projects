package main

import (
	"reflect"
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
