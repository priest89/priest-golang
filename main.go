package main

func main() {
	// var card = "Ace of Spades"
	cards := getNewDeck()
	cards.print()
	cards.shuffle()
	cards.print()
}
