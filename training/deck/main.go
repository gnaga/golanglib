package main



func main() {

	cards := newDeck()
	// hand, rc := deal(cards,3)
	// hand.print()
	// rc.print()
	// cards.saveToFile("my_cards")
	// cards := newDeckFromFile("my_cards")
	cards.shuffle()
	cards.print()

}
