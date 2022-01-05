package main

// All Go programs start with <<package>> keyword, all executable programs MUST have package name <<main>>
import (
	"fmt"
)

func main() {
	myString := fmt.Sprint("Hello World")

	fmt.Printf("%s : My String", myString)
}
