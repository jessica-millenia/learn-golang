package main

import "fmt"

func main() {
	cards := newDeck()
	cards.shuffle()

	cards.saveToFile("my_cards")
	cardsFromFile := newDeckFromFile("my_cards")

	hand1, hand2 := deal(cardsFromFile, 10)
	fmt.Println("Cards in hand 1: ")
	hand1.print()
	fmt.Println("Cards in hand 2: ")
	hand2.print()
}
