package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// Create a new type of 'deck' which is a slice of strings
type deck []string

// create a new deck of cards
func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Clubs", "Diamonds", "Hearts", "Spades"}
	cardValues := []string{"Ace", "2", "3", "4", "5", "6", "7", "8", "9", "10", "Jack", "Queen", "King"}

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

// deal a hand of size x from a deck
// return hand and deck remaining
func deal(d deck, handsize int) (deck, deck) {
	return d[:handsize], d[handsize:]
}

// save to file
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0644)
}

// write deck to string
func (d deck) toString() string {
	return strings.Join([]string(d), ", ")
}
