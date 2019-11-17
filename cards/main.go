package main

import "fmt"

func main() {
	// var card string = "All Cards in a Deck";
	card := "All cards in a deck";
	fmt.Println(card);

	card = newCard();
	fmt.Println(card);
	// card = "updated cards in a deck";
	// fmt.Println(card);
}

func newCard() string {
	return("Five of Dimonds");
}

