package main

import (
	"fmt"

	"github.com/pmarkus/shuffler/deck"
)

func main() {
	fmt.Println("Simulate shuffle")

	d := deck.NewSimpleDeck()
	fmt.Printf("d:\t%s\n", d)

	fmt.Println("== Split ==")
	dTop, dBot := d.Split()
	fmt.Printf("dTop:\t%s\n", dTop)
	fmt.Printf("dBot:\t%s\n", dBot)
	fmt.Printf("d:\t%s\n", d)

	fmt.Println("== Merge ==")
	d = deck.Merge(dTop, dBot)
	fmt.Printf("dTop:\t%s\n", dTop)
	fmt.Printf("dBot:\t%s\n", dBot)
	fmt.Printf("d:\t%s\n", d)
}
