package main

import "fmt"

// Create a new type of 'deck' which is a slice of strings
type deck []string

// create a new deck of cards
func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Clubs", "Diamonds", "Hearts", "Spades"}
	cardValues := []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}

	/*
		When we have a variable that we don't need to use
		you can use the underscore to indicate that the go
		compiler ignore the warning.
	*/
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

// deal a hand of size x from a deck
func deal(d deck, handsize int) (deck, deck) {
	hand := d[:handsize]
	remainingDeck := d[handsize:]
	return hand, remainingDeck
}

// function to print out each card in the deck
/*
	(d deck) is a "receiver" which allows every
	variable of type 'deck' to use the print method.
	It behaves similar to 'this.whatever'.

	The convention is that the receiver variable
	is a one or two letter variable that matches the
	TYPE.
*/
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
