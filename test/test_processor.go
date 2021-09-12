package test

import (
	"github.com/pmarkus/shuffler/deck"
)

type TestProcessor struct {
	iterations []*Iteration
	curIter    *Iteration
}

type Iteration struct {
	startDeck        deck.Deck
	finishDeck       deck.Deck
	posChangeResults []int
}

func NewTestProcessor() *TestProcessor {
	t := TestProcessor{
		iterations: make([]*Iteration, 0),
	}
	return &t
}

func (tp *TestProcessor) StartIteration(d deck.Deck) {
	i := Iteration{
		startDeck:        d,
		posChangeResults: make([]int, len(d.GetCards())),
	}
	tp.iterations = append(tp.iterations, &i)
	tp.curIter = &i
}

func (tp *TestProcessor) FinishIteration(d deck.Deck) {
	tp.curIter.finishDeck = d
}
