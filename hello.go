package main

import "fmt"

const EnglishGreeting = "Hello, "

func Hello (name string) string {
	if name == "" {
		name = "World"
	}
	return EnglishGreeting + name
}

func main () {
	fmt.Println(Hello("John"))
}