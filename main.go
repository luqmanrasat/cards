package main

func main() {
  cards := newDeck()
  // cards := loadDeck("remaining_cards.sav")
  cards.shuffle()
  cards.print()
  // hand, remainingCards := deal(cards, 7)

  // hand.print()
  // remainingCards.saveDeck("remaining_cards.sav")
}
