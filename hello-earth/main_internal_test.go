package main

import "testing"

func ExampleMain() {
	main()
	// Output:
	// Hola!
}

func TestGreet(t *testing.T) {
	type testCase struct {
		lang language
		want string
	}

	var tests = map[string]testCase{
		"English": {
			lang: "en",
			want: "Hello, Earth!",
		},
		"Polish": {
			lang: "pl",
			want: "Dzień dobry!",
		},
		"French": {
			lang: "fr",
			want: "Bonjour!",
		},
		"Spanish": {
			want: "Hola!",
			lang: "es",
		},
		"Urdu": {
			want: `Unsupported language: "ur"`,
			lang: "ur",
		},
		"Empty": {
			lang: "",
			want: `Unsupported language: ""`,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := greet(tc.lang)

			if got != tc.want {
				t.Errorf("expected: %q, got: %q", tc.want, got)
			}
		})
	}

}
