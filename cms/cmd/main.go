package main

import (
	"os"
	"yasir2000/go-web-dev-side-project-1/cms"
)

func main() {
	p := &cms.Page{
		Title:   "Hello, world",
		Content: "this is the body of our webpage",
	}

	cms.Tmpl.ExecuteTemplate(os.Stdout, "index", p)
}
