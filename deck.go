package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// Create a new type of 'deck' which is a slice of strings
type deck []string

// Return a value of type deck
func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Cloves"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nice", "Ten", "Jack", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handsize int) (deck, deck) {
	return d[:handsize], d[handsize:]

}

// Converting our copy of the deck (array slice type) into a string
func (d deck) toString() string {
	return strings.Join([]string(d), ", ")
}

func (d deck) saveToFile(filename string) error {
	// Permission 0666 anyone can read and access to my file
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}
