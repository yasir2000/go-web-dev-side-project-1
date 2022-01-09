package main

type Adder interface {
	Add(a, b int) int
}

type X struct{}

func (x *X) Add(a, b int) int {
	return a + b
}
