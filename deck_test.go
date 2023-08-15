package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}

	if d[0] != "Spades of Ace" {
		t.Errorf("Expect the first card to be Spades of Ace, but got %v", d[0])
	}

	if d[len(d) - 1] != "Clubs of Four" {
		t.Errorf("Expected the last card to be Clubs of Four, but got %v",d[len(d)-1])
	}
}

func TestSaveToFile(t *testing.T) {
	os.Remove("test")
	d := newDeck()
	d.saveToFile("test")

	r := newDeckFromFile("test")

	if len(r) != 16 {
		t.Errorf("Expected the number of cards to be 16, but got %v", len(r))
	}

	os.Remove("test")
}

