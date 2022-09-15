package main

import (
	"fmt"
	
)

func main() {
	cards := []string{"Ace of Diamonds", newCard()}
	cards = append(cards, "Six of Spades")
	
	fmt.Println(cards)
	for i , card := range cards{
		fmt.Println(i, card)
	}
	
	fmt.Println("Iterator")
	for i:=0; i<len(cards); i++{
		fmt.Println(cards[i])
	}
}

func newCard() string {
	return "Five of Diamonds"
}

