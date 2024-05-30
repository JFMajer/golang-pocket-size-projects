package main

import (
	"flag"
	"fmt"
)

type language string

func main() {
	var lang string
	flag.StringVar(&lang, "lang", "en", "Language to print greeting in")
	flag.Parse()

	greeting := greet(language(lang))
	fmt.Println(greeting)
}

var phrasebook = map[language]string{
	"en": "Hello, Earth!",
	"fr": "Bonjour!",
	"pl": "Dzień dobry!",
	"es": "Hola!",
}

func greet(l language) string {
	greeting, ok := phrasebook[l]
	if !ok {
		return fmt.Sprintf("Unsupported language: %q", l)
	}
	return greeting
}
