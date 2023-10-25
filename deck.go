package main

import "fmt"

type deck []string

func newDeck() deck {
  cardSuites := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
  cardValues := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}

  cards := deck{}
  for _, suit := range cardSuites {
    for _, value := range cardValues {
      cards = append(cards, value+" of "+suit)
    }
  }

  return cards
}

func (d deck) print() {
  for _, card := range d {
    fmt.Println(card)
  }
}
