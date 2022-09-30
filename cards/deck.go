package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Hearts", "Clubs", "Diamonds"}

	for _, suit := range cardSuits {
		for i := 1; i < 14; i++ {
			cards = append(cards, strconv.Itoa(i)+" of "+suit)
		}
	}

	return cards
}

func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func byteToDeck(bs []byte) deck {
	return deck(strings.Split(string(bs), ","))
}

func (d deck) deckToByte() []byte {
	return []byte(strings.Join([]string(d), ","))
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, d.deckToByte(), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	return byteToDeck(bs)
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
