package deck

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCuttingDeckLeavesOldDeckEmpty(t *testing.T) {
	d := NewDeck(8)
	d.Cut()
	assert.Equal(t, 0, len(d.cards))
}

func TestCuttingDeckWithOddAmountOfCardsLeaveMoreCardsInBottom(t *testing.T) {
	d := Deck{
		cards: make([]int, 0),
	}
	for i := 0; i < 9; i++ {
		d.cards = append(d.cards, i)
	}
	dTop, dBot := d.Cut()
	assert.True(t, len(dTop.cards) < len(dBot.cards))
}

func TestCuttingDeckWithEvenAmountOfCardsLeaveEqualAmountOfCardsInEachDeck(t *testing.T) {
	d := Deck{
		cards: make([]int, 0),
	}
	for i := 0; i < 8; i++ {
		d.cards = append(d.cards, i)
	}
	dTop, dBot := d.Cut()
	assert.Equal(t, len(dTop.cards), len(dBot.cards))
}

func TestMergingDecksPutFirstDeckOnTop(t *testing.T) {
	d1Cards := []int{
		6,
		5,
	}
	d1 := Deck{
		cards: d1Cards,
	}
	d2Cards := []int{
		4,
		3,
	}
	d2 := Deck{
		cards: d2Cards,
	}
	d := Stack(&d1, &d2)
	assert.Equal(t, d.cards[0], 6)
	assert.Equal(t, d.cards[1], 5)
	assert.Equal(t, d.cards[2], 4)
	assert.Equal(t, d.cards[3], 3)
}
