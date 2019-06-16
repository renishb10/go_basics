package main

import "fmt"

type deck []string

// Func
func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Club"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// This is a reciever
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
