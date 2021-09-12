package main

import (
	"flag"
	"fmt"

	"github.com/pmarkus/shuffler/deck"
	"github.com/pmarkus/shuffler/shuffle"
	"github.com/pmarkus/shuffler/test"
)

func main() {
	runDemo := flag.Bool("demo", false, "Run a small display of functionality.")
	flag.Parse()

	if *runDemo {
		fmt.Println("Demo of a deck and its interactions")
		demo()
	} else {
		simulate()
	}
}

func simulate() {
	deckSize := 99
	shuffleIterations := 10
	simulationIterations := 1000

	fmt.Printf("Deck Size:\t\t%d\n", deckSize)
	fmt.Printf("Shuffle iterations:\t%d\n", shuffleIterations)
	fmt.Printf("Simulation iterations:\t%d\n", simulationIterations)
	fmt.Println("")

	tp := test.NewTestProcessor()

	for i := 0; i < simulationIterations; i++ {

		d := deck.NewSimpleDeck(deckSize)
		tp.StartIteration(*d)

		d = shuffle.TwoPileRotationShuffle(d, shuffleIterations)
		tp.FinishIteration(*d)
	}

	fmt.Println("Position change test: ", test.PositionChangeTest(tp))
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
