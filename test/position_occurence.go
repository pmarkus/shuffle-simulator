package test

import (
	"fmt"
	"math"
)

func PositionOccurenceTest(tp *TestProcessor) string {
	cardCount := len(tp.iterations[0].startDeck.GetCards())
	// realProb := float64(1) / float64(cards)
	expOccur := float64(len(tp.iterations)) / float64(cardCount)
	tolerance := float64(0.15)
	// The occuranceTracker[i] keeps track of how many times card
	// i has been seen in each position in the deck.
	occurenceTracker := make([][]float64, cardCount)
	outsideTolerance := make([]bool, len(occurenceTracker))
	totOutsideTolerance := 0

	for card := 0; card < cardCount; card++ {
		occurenceTracker[card] = positionOccurenceOfCard(card, tp.iterations)
		for position := 0; position < cardCount; position++ {
			occuranceDiff := occurenceTracker[card][position] - expOccur
			// Change the occuranceTracker to instead track the
			// difference from the expected number of occurances.
			// For pretty formatting, round to nearest integer.
			occurenceTracker[card][position] = math.Round(occuranceDiff)

			if math.Abs(occuranceDiff) < expOccur*tolerance {
				// If we are within tolerance, for pretty formatting' sake,
				// Just set the tracker to 0 on that position.
				occurenceTracker[card][position] = 0
			} else {
				outsideTolerance[card] = true
			}
		}
		if outsideTolerance[card] {
			totOutsideTolerance++
		}
	}

	result := fmt.Sprintf("Total cards outside tolerance: %d\n", totOutsideTolerance)
	result += fmt.Sprintf("Real occurence probability at all positions is: %f\n", expOccur)
	result += fmt.Sprintf("Tolerable difference is: +/- %f\n", expOccur*tolerance)
	for i, o := range occurenceTracker {
		if i == 0 {
			result += fmt.Sprintf("Card %d pos:\t[tol=%v]\t", i, !outsideTolerance[i]) + fmt.Sprint(o)
		} else {
			result += fmt.Sprintf("\nCard %d pos:\t[tol=%v]\t", i, !outsideTolerance[i]) + fmt.Sprint(o)
		}
	}
	// result = fmt.Sprintf("Real occurence probability at all positions is: %f\n", realProb) + result
	// result = fmt.Sprintf("Real occurence probability at all positions is: %f\n", expOccur) + result
	return result
}

func positionOccurenceOfCard(card int, iters []*Iteration) []float64 {
	allOccurences := make([]float64, len(iters[0].startDeck.GetCards()))
	for _, iter := range iters {
		cardPos, err := iter.finishDeck.PositionOfCard(card)
		if err != nil {
			panic("a card has been lost in the process")
		}
		allOccurences[cardPos]++
	}
	return allOccurences
}
