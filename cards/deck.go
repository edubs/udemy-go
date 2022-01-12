package main

import (
	crypto_rand "crypto/rand"
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
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
		you can use the _ underscore to indicate that the go
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
	return strings.Join([]string(d), ",")
}

// read string from file
// convert string to slice of strings with split
// cast slice of strings to deck and return it
func readFromFile(filename string) deck {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	s := strings.Split(string(content), ",")
	return deck(s)
}

// shuffle deck
func (d deck) shuffle() {
	// time based seed
	// rand.Seed(time.Now().UnixNano())

	/*
		notes: you should no reseed the random num generator between iterations
		set it once in init or in the function and let it be. the crypto seed
		method below is more secure / reliable than using time.
	*/
	// crypto based seed
	var b [8]byte
	_, err := crypto_rand.Read(b[:])
	if err != nil {
		panic("cannot seed math/rand package with crytpo secure random num")
	}
	rand.Seed(int64(binary.LittleEndian.Uint64(b[:])))

	// get a random index position and swap the values
	// with the current index from loop iteration
	for i := range d {
		randomNumber := rand.Intn(len(d) - 1)
		// long swap method
		/*
			for i, card := range d {
				randCard := d[randomNumber]
				d[randomNumber] = card
				d[i] = randCard
		*/

		// fancy swap
		d[i], d[randomNumber] = d[randomNumber], d[i]
	}
}
