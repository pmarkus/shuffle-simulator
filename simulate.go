package main

import (
	"errors"
	"flag"
	"fmt"

	"github.com/pmarkus/shuffler/deck"
	"github.com/pmarkus/shuffler/shuffle"
)

func main() {
	runDemo := flag.Bool("demo", false, "Run a small display of functionality.")
	quiet := flag.Bool("quiet", true, "only print simulation result")
	flag.Parse()

	if *runDemo {
		fmt.Println("Demo of a deck and its interactions")
		demo()
	} else {
		if !*quiet {
			fmt.Println("Running two-pile rotation-shuffle simulation")
		}
		simulate(*quiet)
	}
}

func simulate(quiet bool) {
	deckSize := 16
	shuffleIterations := 10
	simulationRuns := 1000

	if !quiet {
		fmt.Printf("Deck Size:\t\t" + fmt.Sprint(deckSize) + "\n")
		fmt.Printf("Shuffle iterations:\t" + fmt.Sprint(shuffleIterations) + "\n")
		fmt.Printf("Simulation runs:\t" + fmt.Sprint(simulationRuns) + "\n")
	}

	posTestResults := make([][]int, 0)

	for i := 0; i < simulationRuns; i++ {

		d := deck.NewSimpleDeck(deckSize)
		originalOrder := d.GetCards()
		for i := 0; i < shuffleIterations; i++ {
			d = shuffle.TwoPileRotationShuffle(d, shuffleIterations)
		}

		posTestResults = append(posTestResults, runPositionTest(originalOrder, d.GetCards()))
	}

	averagedPosResults := make([]float64, deckSize)
	averagedPosResultsString := ""
	for i, avg := range averagedPosResults {
		for _, results := range posTestResults {
			avg += float64(results[i])
		}
		avg /= float64(len(posTestResults))
		averagedPosResultsString += fmt.Sprint(avg) + " "
	}

	if !quiet {
		fmt.Println("Simulation results:")
	}
	fmt.Printf("test:\t%s\n", averagedPosResultsString)
}

func runPositionTest(before, after []string) []int {
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

func demo() {
	d := deck.NewSimpleDeck(8)
	fmt.Printf("d:\t%s\n", d)

	fmt.Println("== Split ==")
	dTop, dBot := d.Cut()
	fmt.Printf("dTop:\t%s\n", dTop)
	fmt.Printf("dBot:\t%s\n", dBot)
	fmt.Printf("d:\t%s\n", d)

	fmt.Println("== Stack ==")
	d = deck.Stack(dTop, dBot)
	fmt.Printf("dTop:\t%s\n", dTop)
	fmt.Printf("dBot:\t%s\n", dBot)
	fmt.Printf("d:\t%s\n", d)

	fmt.Println("== Riffle Shuffle ==")
	d.RiffleShuffle()
	fmt.Printf("d:\t%s\n", d)
}
