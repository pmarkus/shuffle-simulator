package shuffle

import "github.com/pmarkus/shuffler/deck"

type Shufflable interface {
	Split() (*deck.Deck, *deck.Deck)
	Merge(*deck.Deck) *deck.Deck
}

func Shuffle(d *deck.Deck) {
	_ = d
}
