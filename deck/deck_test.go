package deck

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSplittingDeckLeavesOldDeckEmpty(t *testing.T) {
	d := NewSimpleDeck()
	d.Split()
	assert.Equal(t, 0, len(d.cards))
}

func TestSplittingDeckWithOddAmountOfCardsLeaveMoreCardsInBottom(t *testing.T) {
	d := Deck{
		cards: make([]card, 0),
	}
	for i := 0; i < 9; i++ {
		d.cards = append(d.cards, card{name: ""})
	}
	dTop, dBot := d.Split()
	assert.True(t, len(dTop.cards) < len(dBot.cards))
}

func TestSplittingDeckWithEvenAmountOfCardsLeaveEqualAmountOfCardsInEachDeck(t *testing.T) {
	d := Deck{
		cards: make([]card, 0),
	}
	for i := 0; i < 8; i++ {
		d.cards = append(d.cards, card{name: ""})
	}
	dTop, dBot := d.Split()
	assert.Equal(t, len(dTop.cards), len(dBot.cards))
}

func TestMergingDecksPutFirstDeckOnTop(t *testing.T) {
	d1Cards := []card{
		{name: "a"},
		{name: "b"},
	}
	d1 := Deck{
		cards: d1Cards,
	}
	d2Cards := []card{
		{name: "c"},
		{name: "d"},
	}
	d2 := Deck{
		cards: d2Cards,
	}
	d := Merge(&d1, &d2)
	assert.Equal(t, d.cards[0].name, "a")
	assert.Equal(t, d.cards[1].name, "b")
	assert.Equal(t, d.cards[2].name, "c")
	assert.Equal(t, d.cards[3].name, "d")
}
