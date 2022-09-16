package main

// import "fmt" 


func main() {
	cards := newDeck()
	// cards.saveToFile("myCards")	

	// cards_copy := newDeckFromFile("myCards")
	// fmt.Println(cards_copy)
	cards.print()
	cards.shuffle()
	cards.print()

}

