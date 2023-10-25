package main

func main() {
  // cards := newDeck()
  cards := loadDeck("remaining_cards.sav")
  hand, remainingCards := deal(cards, 7)

  hand.print()
  remainingCards.saveToFile("remaining_cards.sav")
}
