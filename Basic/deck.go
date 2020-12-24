package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func newDack() deck {
	cards := deck{}
	cardSuites := []string{"Spades", "Dimonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight",
		"Nine", "Ten", "Jocker", "Queen", "King"}

	for _, suite := range cardSuites {
		for _, value := range cardValues {
			cards = append(cards, suite+" of "+value)
		}
	}

	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func (d deck) toString() string {
	str := strings.Join(d, ",")
	return str
}

func (d deck) saveToFile(filename string) error {
	error := ioutil.WriteFile(filename, []byte(d.toString()), 0666)
	return error
}

func (d deck) shuffle() {
	// Make sure to use new seed for rand type
	// New source requires Int64 so we are passing UnixNano()
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	length := len(d)

	// Not the best logic but it will make the job done
	for idx := range d {
		newIdx := r.Intn(length)
		d[idx], d[newIdx] = d[newIdx], d[idx]
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("ERROR: ", err)
		os.Exit(1)
	}
	return strings.Split(string(bs), ",")
}
