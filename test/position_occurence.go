package test

import "fmt"

func PositionOccurenceTest(tp *TestProcessor) string {
	occurenceTracker := make([][]float64, len(tp.iterations[0].startDeck.GetCards()))
	for i := range occurenceTracker {
		occurenceTracker[i] = positionOccurenceOfCard(i, tp.iterations)
		for j := range occurenceTracker {
			occurenceTracker[i][j] /= float64(len(tp.iterations))
		}
	}
	result := ""
	for i, o := range occurenceTracker {
		if i == 0 {
			result += fmt.Sprintf("Card %d pos:\t", i) + fmt.Sprint(o)
		} else {
			result += fmt.Sprintf("\nCard %d pos:\t", i) + fmt.Sprint(o)
		}
	}
	var optimalOccurence float64 = 1.0 / float64(len(tp.iterations[0].startDeck.GetCards()))
	result = fmt.Sprintf("Optimal occurence at all positions is: %f\n", optimalOccurence) + result
	return result
}

func positionOccurenceOfCard(card int, iters []*Iteration) []float64 {
	allOccurences := make([]float64, len(iters[0].startDeck.GetCards()))
	for _, iter := range iters {
		cardPos, err := iter.finishDeck.PositionOfCard(card)
		if err != nil {
			panic("a card have been lost in the process")
		}
		allOccurences[cardPos]++
	}
	return allOccurences
}
