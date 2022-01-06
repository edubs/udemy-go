package main

import "fmt"

func main() {
	cards := newDeck()

	// deal returns two values of type 'deck'
	hand, remainingDeck := deal(cards, 5)

	hand.print()
	remainingDeck.print()

	fmt.Println(hand.toString())
	fmt.Println(toBytes(hand.toString()))

}
