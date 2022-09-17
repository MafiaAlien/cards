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

}


func TestSaveToDeckAndNewDeckFromFile(t *testing.T){
	os.Remove("_decktesting")
	
	d := newDeck()
	d.saveToFile("_decktesting")
	
	d_copy := newDeckFromFile("_decktesting")
	if len(d_copy) != 16 {
		t.Errorf("Expected 16 cards in deck, got %v", len(d_copy))
	}

	os.Remove("_decktesting")
}