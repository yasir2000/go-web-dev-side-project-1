package main

import (
	"fmt"
	"net/http"
	"time"
)

var sites = []string{
	"https://github.com",
	"https://google.com",
	"https://stackoverflow.com",
	"https://facebook.com",
	"https://twitter.com",
	"https://golang.org",
	"https://forum.golangbridge.org",
	"https://packtpub.com",
}

func get() {
	start := time.Now()
	for _, s := range sites {
		res, _ := http.Get(s)
		fmt.Printf("%s %d\n", s, res.StatusCode)
		res.Body.Close()
		elapsed := time.Since(start)
		fmt.Println(elapsed)
	}
}

func getConcurrently() {
	start := time.Now()
	ch := make(chan string)

	for _, s := range sites {
		go func(s string) {
			res, _ := http.Get(s)
			ch <- fmt.Sprintf("%s %d", s, res.StatusCode)
		}(s)
	}

	for range sites {
		fmt.Println(<-ch)

		elapsed := time.Since(start)
		fmt.Println(elapsed)
	}

}

// func main() {
// 	switch os.Args[1] {
// 	case "seq":
// 		get()
// 	case "con":
// 		getConcurrently()
// 	default:
// 		fmt.Println("Please choose `seq` or `con` .")
// 	}
// }
