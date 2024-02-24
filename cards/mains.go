package main

import "fmt"

func main() {
	cards := deck{"Ace of diamonds", "Spade of diamonds"}
	fmt.Println(cards)

	cards.print()
}
