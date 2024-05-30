package main

import "fmt"

type language string

func main() {
	greeting := greet("en")
	fmt.Println(greeting)
}

func greet(l language) string {
	switch l {
	case "en":
		return "Hello, Earth!"
	case "fr":
		return "Bonjour!"
	case "pl":
		return "Dzień dobry!"
	default:
		return ""
	}

}
