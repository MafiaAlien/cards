package main

import "fmt"

// create a new type of deck
// which is a slice of strings

type Deck []string

func (d Deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

func newDeck() Deck {
	cards := Deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, suit+" of "+value)
		}
	}
	return cards

}

func (d Deck) shuffle() {

}
