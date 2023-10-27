package main

import (
  "testing"
  "os"
)

func TestNewDeck(t *testing.T) {
  d := newDeck()

  if len(d) != 52 {
    t.Errorf("Expected deck length of 52, but got %v", len(d))
  }

  if d[0] != "1 of ♠️" {
    t.Errorf("Expected deck starts with 1 of ♠️, but got %v", d[0])
  }

  if d[len(d)-1] != "K of ♦️" {
    t.Errorf("Expected deck ends with K of ♦️, but got %v", d[len(d)-1])
  }
}

func TestSaveDeckAndLoadDeck(t *testing.T) {
  filename := "_decktesting"
  os.Remove(filename)

  d := newDeck()
  d.saveDeck(filename)

  loadedDeck := loadDeck(filename)

  if len(loadedDeck) != 52 {
    t.Errorf("Expected loaded deck length of 52, but got %v", len(loadedDeck))
  }
  
  os.Remove(filename)
}
