package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	env := strings.Split(os.Getenv("PATH"), ";")
	//fmt.Println(env)

	for _, val := range env {
		fmt.Printf("%s\n", val)
	}
}
