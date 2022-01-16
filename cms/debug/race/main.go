// This example triggers warnings from the race detector.
//
// Mutates a global value at random. Output is different everytime the program
// is run. Illustrates a (very basic) data race.
package main

import (
	"fmt"
	"math/rand"
	"time"
)

// N is our gloabl mutable value
var N = 1

func main() {
	// Try commenting an un-commenting this line of code!
	//
	// With this line *commented* you'll get deterministic output, since the
	// seed is set for the math/rand package when your machine starts up.
	//
	// With this *uncommented* you'll random output, since the seed is set
	// each time the program runs.
	//
	// The race detector is smart enough to detect a data race in both cases.
	rand.Seed(time.Now().UnixNano())

	go mutateAtRandom()
	mutateAtRandom()

	fmt.Println(N == 1, N)
}

func mutateAtRandom() {
	v := time.Duration(rand.Intn(3))
	time.Sleep(v * time.Second)
	N = int(v)
}
