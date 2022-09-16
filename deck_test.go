package main 

import (
	"testing"
)


func TestNewDeck(t *testing.T){
	cards := newDeck()
	cards.print()
}