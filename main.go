package main

import "fmt"

func main() {
	// var card = "Ace of Spades"
	cards := getNewDeck()
	hand, remainingCards := deal(cards, 5)
	hand.print()
	remainingCards.print()
	fmt.Println(hand.toString())
}
