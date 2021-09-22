package main

import "fmt"

func main() {
	cards := newDeck()
	cards.shuffle()
	cards.print()
	newDeck, oldDeck := deal(cards, 5)
	fmt.Println("\nNew deck")
	newDeck.print()
	fmt.Println("\nOld deck")
	oldDeck.print()
}