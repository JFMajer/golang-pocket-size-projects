package main

import "testing"

func ExampleMain() {
	main()
	// Output:
	// Hello, Earth!
}

func TestGreet_English(t *testing.T) {
	lang := language("en")
	want := "Hello, Earth!"

	got := greet(lang)

	if want != got {
		t.Errorf("expected: %q, got: %q", want, got)
	}
}

func TestGreet_Polish(t *testing.T) {
	lang := language("pl")
	want := "Dzień dobry!"

	got := greet(lang)

	if want != got {
		t.Errorf("expected: %q, got: %q", want, got)
	}
}

func TestGreet_French(t *testing.T) {
	lang := language("fr")
	want := "Bonjour!"

	got := greet(lang)

	if want != got {
		t.Errorf("expected: %q, got: %q", want, got)
	}
}
