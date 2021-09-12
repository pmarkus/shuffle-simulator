package deck

import (
	"strconv"
)

type Deck struct {
	cards []card
}

type card struct {
	name string
}

func NewSimpleDeck() *Deck {
	var size int = 8
	cards := make([]card, 0)
	for i := 0; i < size; i++ {
		c := card{
			name: strconv.Itoa(i),
		}
		cards = append(cards, c)
	}

	deck := Deck{
		cards: cards,
	}
	return &deck
}

func (d *Deck) Split() (*Deck, *Deck) {
	dTop := Deck{
		cards: make([]card, 0),
	}
	dBot := Deck{
		cards: make([]card, 0),
	}
	var half int = len(d.cards) / 2
	for i := 0; i < half; i++ {
		dTop.cards = append(dTop.cards, d.cards[i])
	}
	for i := half; i < len(d.cards); i++ {
		dBot.cards = append(dBot.cards, d.cards[i])
	}
	d.cards = make([]card, 0)
	return &dTop, &dBot
}

func Merge(dTop *Deck, dBot *Deck) *Deck {
	d := Deck{
		cards: make([]card, 0),
	}
	d.cards = append(d.cards, dTop.cards...)
	d.cards = append(d.cards, dBot.cards...)
	dTop.cards = make([]card, 0)
	dBot.cards = make([]card, 0)

	return &d
}

func (d *Deck) String() string {
	s := ""
	for i, card := range d.cards {
		if i != 0 {
			s = s + "\t" + card.name
		} else {
			s = s + card.name
		}
	}
	return s
}
