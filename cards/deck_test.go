package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(d))
	}
	if d[0] != "1 of Spades" {
		t.Errorf("Expected first card of 1 of Spades, but got %v", d[0])
	}
	if d[len(d)-1] != "13 of Diamonds" {
		t.Errorf("Expected last card of 13 of Diamonds, but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 52 {
		t.Errorf("Expected 52 cards in deck, got %v", len(loadedDeck))
	}
	if loadedDeck[0] != "1 of Spades" {
		t.Errorf("Expected first card of 1 of Spades, but got %v", loadedDeck[0])
	}
	if loadedDeck[len(loadedDeck)-1] != "13 of Diamonds" {
		t.Errorf("Expected last card of 13 of Diamonds, but got %v", loadedDeck[len(loadedDeck)-1])
	}

	os.Remove("_decktesting")
}
