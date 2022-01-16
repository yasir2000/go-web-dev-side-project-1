// This example illustrates type reflection.
//
// There are two examples: a simple one to demonstrate the strange behaviour,
// and another one to highlight a useful, real-world example.
package main

import (
	"fmt"
	"reflect"
)

type a struct {
	B string
	C int
}

func main() {
	x := &a{
		B: "B",
		C: 1,
	}
	y := &a{
		B: "B",
		C: 1,
	}

	fmt.Println(x == y)
	fmt.Println(&x, &y)
	// Using deep equal gives us the results we want:
	fmt.Println(reflect.DeepEqual(x, y))
}
