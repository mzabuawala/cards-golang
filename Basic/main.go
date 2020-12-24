package main

import "fmt"

func main() {
	handsFile := "hands.txt"
	remainingCardsFile := "remaingCards.txt"
	// New Deck of cards
	cards := newDack()

	fmt.Println("Before shuffle: ", cards)
	// Shuffle the cards
	cards.shuffle()
	fmt.Println("After shuffle: ", cards)

	// Pass the Hand
	hand, remainingCards := deal(cards, 5)

	// Save it in the file
	hand.saveToFile(handsFile)
	remainingCards.saveToFile(remainingCardsFile)

	// Read the file
	newHand := newDeckFromFile(handsFile)
	fmt.Println(newHand)

	// Testing out an errr handling code
	newHand = newDeckFromFile("DoNotExists.txt")
	fmt.Println(newHand)

}
