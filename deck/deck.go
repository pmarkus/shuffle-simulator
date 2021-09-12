package deck

import (
	"errors"
	"math/rand"
	"strconv"
	"time"
)

type Deck struct {
	cards []int
}

func NewSimpleDeck(size int) *Deck {
	cards := make([]int, size)
	for i := 0; i < size; i++ {
		cards[i] = i
	}
	deck := Deck{
		cards: cards,
	}
	return &deck
}

func (d *Deck) GetCard(i int) (int, error) {
	if i < 0 || i >= len(d.cards) {
		return -1, errors.New("index out of bounds")
	}
	return d.cards[i], nil
}

func (d *Deck) GetCards() []int {
	return d.cards
}

func (d *Deck) Cut() (*Deck, *Deck) {
	dTop := Deck{
		cards: make([]int, 0),
	}
	dBot := Deck{
		cards: make([]int, 0),
	}
	var half int = len(d.cards) / 2
	for i := 0; i < half; i++ {
		dTop.cards = append(dTop.cards, d.cards[i])
	}
	for i := half; i < len(d.cards); i++ {
		dBot.cards = append(dBot.cards, d.cards[i])
	}
	d.cards = make([]int, 0)
	return &dTop, &dBot
}

func Stack(dTop *Deck, dBot *Deck) *Deck {
	d := Deck{
		cards: make([]int, 0),
	}
	d.cards = append(d.cards, dTop.cards...)
	d.cards = append(d.cards, dBot.cards...)
	dTop.cards = make([]int, 0)
	dBot.cards = make([]int, 0)

	return &d
}

func (d *Deck) RiffleShuffle() {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	dSize := len(d.cards)
	p1, p2 := d.Cut()
	p1Size := len(p1.cards)
	p2Size := len(p2.cards)
	d.cards = make([]int, 0, dSize)

	for len(d.cards) < dSize {
		var probP1 float64 = float64(p1Size) / (float64(p1Size) + float64(p2Size))
		if r.Float64() < probP1 {
			d.cards = append(d.cards, p1.cards[len(p1.cards)-p1Size])
			p1Size--
		} else {
			d.cards = append(d.cards, p2.cards[len(p2.cards)-p2Size])
			p2Size--
		}
	}
	p1.cards = make([]int, 0)
	p2.cards = make([]int, 0)
}

func (d *Deck) PositionOfCard(card int) (int, error) {
	for i, c := range d.cards {
		if c == card {
			return i, nil
		}
	}
	return -1, errors.New("card not found")
}

func (d *Deck) String() string {
	s := ""
	for i, card := range d.cards {
		if i != 0 {
			s = s + "\t" + strconv.Itoa(card)
		} else {
			s = s + strconv.Itoa(card)
		}
	}
	return s
}
