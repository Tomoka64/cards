package main

import (
	"fmt"

	"github.com/fatih/color"
)

func main() {
	cards := newDeck()
	cards.shuffle()
	deck1, yourcards := deal(cards, len(cards)/2)
	deck1.SaveToFile("cards.txt")
	fmt.Println("out of 16 cards you got...")
	yourcards.print()
	color.Red("after shuffle:\n")
	card := split(yourcards)
	removeDuplicates(card)
}
