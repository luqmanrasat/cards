package main

import (
  "fmt"
  "strings"
  "io/ioutil"
  "os"
  "math/rand"
  "time"
)

type deck []string

func newDeck() deck {
  cardSuites := []string{"♠️", "♥️", "♣️", "♦️"}
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
  for i, card := range d {
    fmt.Println(i, card)
  }
}

func deal(d deck, handSize int) (deck, deck) {
  return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
  return strings.Join([]string(d), ",")
}

func (d deck) saveDeck(filename string) error {
  return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func loadDeck(filename string) deck {
  bs, err := ioutil.ReadFile(filename)

  if err != nil {
    fmt.Println("Error:", err)
    os.Exit(1)
  }

  s := strings.Split(string(bs), ",")
  return deck(s)
}

func (d deck) shuffle() {
  source :=  rand.NewSource(time.Now().UnixNano())
  r := rand.New(source)

  for i := range d {
    newPosition := r.Intn(len(d) - 1)
    d[i], d[newPosition] = d[newPosition], d[i]
  }
}
