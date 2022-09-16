package main

import (
	"time"
	"fmt"
	"os"
	"strings"
	"math/rand"
	
)

// create a new type of deck
// which is a slice of strings

type Deck []string

func (d Deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d Deck, handSize int) (Deck, Deck) {
	return d[:handSize], d[handSize:]
}

func newDeck() Deck {
	cards := Deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards

}

func (d Deck) toString() string {
		return strings.Join([]string(d),",")
	}	

func (d Deck) saveToFile(filename string) error {
		return os.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) Deck {
	 bs, err :=  os.ReadFile(filename)
	 if err != nil {
		// log the error
		fmt.Println("Error: ", err)
		// log the error and quit the program entirely 
		os.Exit(1)
	 } 
	 
	 s :=  strings.Split(string(bs), ",")
	 
	 return Deck(s)
	}

func (d Deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	
	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
		}
	}
	
