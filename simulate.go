package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/pmarkus/shuffler/deck"
	"github.com/pmarkus/shuffler/shuffle"
	"github.com/pmarkus/shuffler/test"
)

func main() {
	runDemo := flag.Bool("demo", false, "Run a small display of functionality.")
	hardRandom := flag.Bool(
		"hard",
		false,
		"Include a \"hard random\" (near perfect random) result at the end. Defaults to false.",
	)
	onlyCustom := flag.Bool(
		"custom",
		false,
		"Show only result from the custom shuffling scheme. Defaults to false.",
	)
	deckSize := flag.Int(
		"cards",
		52,
		"The amount of cards in the deck to be shuffled. Defaults to 52.",
	)
	shuffleTimes := flag.Int(
		"shuffles",
		9,
		"The amount of times to shuffle the deck for each shuffling scheme. Defaults to 9.",
	)
	simulationIterations := flag.Int(
		"iterations",
		100_000,
		"The amount of iterations to run the simulation before gathering results. Defaults to 100,000.",
	)
	flag.Parse()

	if *runDemo {
		fmt.Println("Demo of a deck and its interactions")
		demo()
	} else {
		simulate(*deckSize, *shuffleTimes, *simulationIterations, *onlyCustom, *hardRandom)
	}
}

func simulate(deckSize, shuffleTimes, simulationIterations int, onlyCustomTest, hardRandomTest bool) {
	startT := time.Now().UnixNano()
	fmt.Printf("Deck Size:\t\t%d\n", deckSize)
	fmt.Printf("Shuffle iterations:\t%d\n", shuffleTimes)
	fmt.Printf("Simulation iterations:\t%d\n", simulationIterations)
	fmt.Println("")

	// Custom shuffling scheme test
	tp := test.NewTestProcessor()
	// Riffle shuffle test
	rstp := test.NewTestProcessor()
	// Riffle shuffle test with cutting - swapping top and bottom halves - halfway through.
	rswctp := test.NewTestProcessor()
	// Hard random test. I.e. as close to perfect random as we can get.
	hrtp := test.NewTestProcessor()

	for i := 0; i < simulationIterations; i++ {
		// Custom algoritm
		d := deck.NewDeck(deckSize)
		tp.StartIteration(*d)

		d = shuffle.TwoPileRotationShuffle(d, shuffleTimes)
		tp.FinishIteration(*d)

		// Normal riffle shuffle on whole deck
		if !onlyCustomTest {
			rsd := deck.NewDeck(deckSize)
			rstp.StartIteration(*rsd)

			for i := 0; i < shuffleTimes; i++ {
				rsd.RiffleShuffle()
			}
			rstp.FinishIteration(*rsd)
		}

		if !onlyCustomTest {
			// Riffle shuffle deck, cutting halfway through all shuffles
			rswcd := deck.NewDeck(deckSize)
			rswctp.StartIteration(*rswcd)

			for i := 0; i < shuffleTimes/2; i++ {
				rswcd.RiffleShuffle()
			}
			top, bot := rswcd.Cut()
			rswcd = deck.Stack(bot, top)
			for i := int(shuffleTimes / 2); i < shuffleTimes; i++ {
				rswcd.RiffleShuffle()
			}
			rswctp.FinishIteration(*rswcd)
		}

		// Hard random deck
		if hardRandomTest {
			hrd := deck.NewDeck(deckSize)
			hrtp.StartIteration(*hrd)
			hrd = deck.NewRandomDeck(deckSize)
			hrtp.FinishIteration(*hrd)
		}
	}

	// Custom algoritm
	fmt.Printf("\nCustom shuffled deck\n")
	fmt.Printf("Position change test:\n%s\n\n", test.PositionChangeTest(tp))
	fmt.Printf("Position occurence test:\n%s\n", test.PositionOccurenceTest(tp))

	// Normal riffle shuffle on whole deck
	if !onlyCustomTest {
		fmt.Printf("\nNormal riffle shuffled deck\n")
		fmt.Printf("Position change test:\n%s\n\n", test.PositionChangeTest(rstp))
		fmt.Printf("Position occurence test:\n%s\n", test.PositionOccurenceTest(rstp))
	}

	// Riffle shuffle deck, cutting halfway through all shuffles
	if !onlyCustomTest {
		fmt.Printf("\nRiffle shuffled deck, with a cut halfway through\n")
		fmt.Printf("Position change test:\n%s\n\n", test.PositionChangeTest(rswctp))
		fmt.Printf("Position occurence test:\n%s\n", test.PositionOccurenceTest(rswctp))
	}

	if hardRandomTest {
		// Hard random deck
		fmt.Printf("\nHard random deck\n")
		fmt.Printf("Position change test:\n%s\n\n", test.PositionChangeTest(hrtp))
		fmt.Printf("Position occurence test:\n%s\n", test.PositionOccurenceTest(hrtp))
	}

	endT := time.Now().UnixNano()
	fmt.Println("")
	fmt.Printf("Elapsed milliseconds: %v\n", ((endT - startT) / 1000000))
}

func demo() {
	d := deck.NewDeck(8)
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
