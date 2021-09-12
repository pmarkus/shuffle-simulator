package shuffle

import "github.com/pmarkus/shuffler/deck"

func TwoPileRotationShuffle(d *deck.Deck, t int) *deck.Deck {
	p1, p2 := d.Cut()
	for i := 0; i < t; i++ {
		p1.RiffleShuffle()
		p2.RiffleShuffle()
		p1Top, p1Bot := p1.Cut()
		p2Top, p2Bot := p2.Cut()
		p1 = deck.Stack(p1Top, p2Bot)
		p2 = deck.Stack(p2Top, p1Bot)
	}
	return deck.Stack(p1, p2)
}
