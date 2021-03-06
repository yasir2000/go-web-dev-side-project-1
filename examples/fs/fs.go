package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

func main() {
	var dir string
	port := flag.String("port", "3000", "port to serve ftp on")
	path := flag.String("path", "", "path to serve ftp")
	flag.Parse()

	if *path == "" {
		dir, _ = os.Getwd()
	} else {
		dir = *path
	}
	log.Fatal(http.ListenAndServe(":"+*port, http.FileServer(http.Dir(dir))))
}
