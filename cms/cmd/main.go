package main

import (
	"net/http"
	"os"
	"yasir2000/go-web-dev-side-project-1/cms"
)

func main() {
	// p := &cms.Page{
	// 	Title:   "Hello, world",
	// 	Content: "this is the body of our webpage",
	// }

	// cms.Tmpl.ExecuteTemplate(os.Stdout, "index", p)

	http.HandleFunc("/", cms.ServeIndex)
	http.HandleFunc("/new", cms.HandleNew)

	if os.Getenv("env") == "dev" {
		http.ListenAndServe(":3000", nil)
	} else {
		http.ListenAndServe(":80", nil)
	}
}
