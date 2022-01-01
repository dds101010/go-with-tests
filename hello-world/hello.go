package main

import "fmt"

const EnglishGreeting = "Hello, "
const SpanishGreeting = "Hola, "
const FrenchGreeting = "Bonjour, "

func Hello (name string, language string) string {
	if name == "" {
		name = "World"
	}
	return greetingPrefix(language) + name
}

func greetingPrefix (language string) (prefix string) {
	switch language {
	case "French":
		prefix = FrenchGreeting
	case "Spanish":
		prefix = SpanishGreeting
	default:
		prefix = EnglishGreeting
	}
	return
}

func main () {
	fmt.Println(Hello("John", "English"))
}