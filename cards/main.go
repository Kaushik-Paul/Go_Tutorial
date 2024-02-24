package main

import "fmt"

func main() {
	// First way of assigning variable
	var card string = "Ace of Spades"
	fmt.Println(card)

	// Second way of assigning variable
	newCard := "New Ace of Spade"
	fmt.Println(newCard)

	newCard = "Five of Diamonds"
	fmt.Println(newCard)

	newCard = newCardFunc()
	fmt.Println("Returned String: " + newCard)

	// From separate file
	printState()

	// Arrays and slice
	cards := []string{newCardFunc(), "Ace of diamonds"}
	cards = append(cards, "Six of spades")

	fmt.Println(cards)

	// For loops

}

func newCardFunc() string {
	return "Five of Diamonds"
}
