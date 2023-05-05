package main

func main() {
	// Creating a new deck and save to file
	//	cards := newDeck()
	//	cards.saveToFile("my_cards")

	// Loading a deck from a file
	//cards := newDeckFromFile("my")
	//cards.print()

	// Create new deck and shuffle
	cards := newDeck()
	cards.shuffle()
	cards.print()
}
