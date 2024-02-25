package main

import "fmt"

func main() {
	cards := newDeck()
	cards.print()
	fmt.Println("==================================================")

	fmt.Println(cards.toString())
}
