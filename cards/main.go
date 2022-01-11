package main

func main() {
	cards := readFromFile("my-cards.txt")
	cards.shuffle()
	cards.print()

}
