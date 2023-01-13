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

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "King of Clubs" {
		t.Errorf("Expected last card of King of Clubs, but got %v", d[51])
	}
}

func TestSaveToDeckAndNewDeckTestFromFile(t *testing.T) {
	//remove any old files
	os.Remove("_decktesting")
	d := newDeck()

	d.saveToFile("_decktesting")

	ld := newDeckFromFile("_decktesting")

	if len(ld) != 52 {
		t.Errorf("Expected 52 cards in deck, but got %v", len(ld))
	}

	os.Remove("_decktesting")
}
