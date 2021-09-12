package test

import (
	"errors"
	"fmt"
)

func PositionChangeTest(tp *TestProcessor) string {
	totalPosChanges := make([]int, len(tp.iterations[0].startDeck.GetCards()))
	for _, iter := range tp.iterations {
		posChanges := iterationPosChanges(iter.startDeck.GetCards(), iter.finishDeck.GetCards())
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

func iterationPosChanges(before, after []string) []int {
	r := make([]int, 0)
	for iBefore, bName := range before {
		iAfter, err := findPosOfCard(after, bName)
		if err != nil {
			panic("a card have been lost in the process")
		}
		r = append(r, iBefore-iAfter)
	}
	return r
}

func findPosOfCard(s []string, name string) (int, error) {
	for i, sName := range s {
		if sName == name {
			return i, nil
		}
	}
	return -1, errors.New("card not found")
}
