package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := getNewDeck()
	if len(d) != 16 {
		t.Errorf("Expected length of deck is 16 %v", len(d))
	}

	if d[0] != "Ace of Diamonds" {
		t.Errorf("Expected get Ace of Diamond, but got %v", d[0])
	}

	if d[len(d)-1] != "Four of Spades" {
		t.Errorf("Expected get Four of Spades, but got %v", d[len(d)-1])
	}
}

func TestSaveDeckToFileAndLoadNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	newDeck := getNewDeck()
	newDeck.writeToFile("_decktesting")
	loadedDeck := getDeckFromFile("_decktesting")

	if len(loadedDeck) != 16 {
		t.Errorf("Can not loaded deck from file %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
