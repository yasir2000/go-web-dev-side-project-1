// This example shows how to get information with delve.
package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type user struct {
	Name string
	Age  int
}

func main() {
	r := strings.NewReader(`
		{
			"Name": "Ben",
			"Age":  100
		}
	`)
	var v user

	// This should be `json.NewDecoder(r).Decode(&v)`, but it can be hard
	// to catch, especially if you're new and ignore errors!
	json.NewDecoder(r).Decode(&v)

	fmt.Printf("%v\n", v)
}
