package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct {
}

type spanishBot struct {
}

func main() {
	/*
		this is how you declare a new a new value of type struct
		when the struct has no fields associated with it
	*/
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

/*
	if you are a type that has a function called getGreeting()
	that returns a string, then you are also of type 'bot' via
	the intereface. this allows both eb and sb to call printGreeting
	even though they are not directly of type 'bot'
*/
func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	return "Hi there!"
}

func (spanishBot) getGreeting() string {
	return "Hola amigo!"
}
