package main

func main() {
	cards := newDeck()
	// hand, remainingDeck := deal(cards, 5)
	// hand.print()
	// remainingDeck.print()
	// cards.print()
	// fmt.Println(cards.toString())
	// cards.saveToFile("my_cards")

	//newDeckFromFile("my_cards").print()
	//newDeckFromFile("my_cards2").print()
	cards.shuffle()
	cards.print()
}
