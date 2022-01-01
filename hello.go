package main

import "fmt"

const EnglishGreeting = "Hello, "
const SpanishGreeting = "Hola, "

func Hello (name string, language string) string {
	if name == "" {
		name = "World"
	}
	if language == "Spanish" {
		return SpanishGreeting + name
	}
	return EnglishGreeting + name
}

func main () {
	fmt.Println(Hello("John", "English"))
}