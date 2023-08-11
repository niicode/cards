package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

type deck []string

func newDeck () deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string {"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
	  for _, value := range cardValues {
		cards = append(cards, suit+ " of "+value,)
	  }
	}
	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int)(deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func toDeck(s []byte) deck {
	d := strings.Split(string(s), ",")

	return deck(d)

}

func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := os.ReadFile(filename)

	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}

	return toDeck(bs)
}

func (d deck) shuffle() {
	for i := range d {
		newPos := rand.Intn(len(d) - 1)
		d[i], d[newPos] = d[newPos], d[i]
	}
}
