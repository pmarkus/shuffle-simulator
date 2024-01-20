package deck

import (
	"errors"
	"math"
	"math/rand"
	"strconv"
	"time"
)

var (
	rndSeed = rand.NewSource(time.Now().UnixNano())
	rnd     = rand.New(rndSeed)
)

type Deck struct {
	cards []int
}

func NewDeck(size int) *Deck {
	cards := make([]int, size)
	for i := 0; i < size; i++ {
		cards[i] = i
	}
	deck := Deck{
		cards: cards,
	}
	return &deck
}

func NewRandomDeck(size int) *Deck {
	cards := make([]int, size)
	seq := make([]int, size)
	for i := range seq {
		seq[i] = i
	}

	for i := 0; i < size; i++ {
		index := rnd.Intn(len(seq))
		cards[i] = seq[index]
		seq = remove(seq, index)
	}
	deck := Deck{
		cards: cards,
	}
	return &deck
}

func remove(s []int, i int) []int {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
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

	cutIndex := getCutPointIndexByNormalDistribution(len(d.cards))
	for i := 0; i < cutIndex; i++ {
		dTop.cards = append(dTop.cards, d.cards[i])
	}
	for i := cutIndex; i < len(d.cards); i++ {
		dBot.cards = append(dBot.cards, d.cards[i])
	}
	d.cards = make([]int, 0)
	return &dTop, &dBot
}

// Calculates a random index close to the half of the deck according to a normal distribution.
// This simulates a human's atempt at cutting a deck in half.
// The returned index will always be in the range (0, deckSize-1].
func getCutPointIndexByNormalDistribution(deckSize int) int {
	standardDev := float64(deckSize) / 20
	// The half-point compensation is to compensate for us working with a set of numbers (indices).
	// Like how the average value of a six-sided dice [1,2,3,4,5,6] isn't 3, but 3.5.
	const halfPointCompensation = 0.5

	index := int(math.Round(rnd.NormFloat64()*standardDev + float64(deckSize-1)/2 + halfPointCompensation))
	// Do not allow the cut index to be the first element. I.e. cut index must be at least index 1.
	if index < 1 {
		index = 1
	}
	if index > deckSize-1 {
		index = deckSize - 1
	}

	return index
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
	dSize := len(d.cards)
	p1, p2 := d.Cut()
	p1Size := len(p1.cards)
	p2Size := len(p2.cards)
	d.cards = make([]int, 0, dSize)

	for len(d.cards) < dSize {
		var probP1 float64 = float64(p1Size) / (float64(p1Size) + float64(p2Size))
		if rnd.Float64() < probP1 {
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
