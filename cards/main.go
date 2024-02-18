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
}
