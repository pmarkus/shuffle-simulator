package shuffle

import "github.com/pmarkus/shuffler/deck"

func TwoPileRotationShuffle(d *deck.Deck, shuffleCount int) *deck.Deck {
	p1, p2 := d.Cut()
	for i := 0; i < shuffleCount; i++ {
		p1.RiffleShuffle()
		p2.RiffleShuffle()
		p1Top, p1Bot := p1.Cut()
		p2Top, p2Bot := p2.Cut()
		// p1 = deck.Stack(p1Top, p2Bot)
		// p2 = deck.Stack(p2Top, p1Bot)
		// The scheme is slightly more efficient if exchanging places of top with bottom halves.
		p1 = deck.Stack(p2Bot, p1Top)
		p2 = deck.Stack(p1Bot, p2Top)
	}
	return deck.Stack(p1, p2)
}
