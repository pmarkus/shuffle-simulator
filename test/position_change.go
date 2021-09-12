package test

import (
	"errors"
	"fmt"

	"github.com/pmarkus/shuffler/deck"
)

func PositionChangeTest(tp *TestProcessor) string {
	totalPosChanges := make([]int, len(tp.iterations[0].startDeck.GetCards()))
	for _, iter := range tp.iterations {
		posChanges := iterationPosChanges(iter.startDeck, iter.finishDeck)
		for i, change := range posChanges {
			totalPosChanges[i] += change
		}
	}
	avaragePosChanges := make([]float64, len(totalPosChanges))
	for i, c := range totalPosChanges {
		avaragePosChanges[i] = float64(c) / float64(len(tp.iterations))
	}

	return fmt.Sprint(avaragePosChanges)
}

func iterationPosChanges(before, after deck.Deck) []int {
	r := make([]int, 0)
	for iBefore, card := range before.GetCards() {
		iAfter, err := after.PositionOfCard(card)
		if err != nil {
			panic("a card have been lost in the process")
		}
		r = append(r, iBefore-iAfter)
	}
	return r
}

func findPosOfCard(cards []int, card int) (int, error) {
	for i, cName := range cards {
		if cName == card {
			return i, nil
		}
	}
	return -1, errors.New("card not found")
}
