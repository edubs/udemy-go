package main

func main() {
	cards := newDeck()

	// deal returns two values of type 'deck'
	hand, remainingDeck := deal(cards, 5)

	hand.print()
	remainingDeck.print()

}
