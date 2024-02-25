package main

func main() {
	//cards := newDeck()
	//cards.print()
	//fmt.Println("==================================================")
	//
	//fmt.Println(cards.toString())

	//cards.saveToFile("my_cards")
	cards := newDeckFromFile("my_cards")
	cards.print()
}
