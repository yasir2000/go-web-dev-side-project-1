package cms

import (
	"html/template"
	"os"
	"time"
)

var temlPath = os.Getenv("GOPATH") + "src/yasir2000/go-web-dev-side-project-1/cms")

//Tmpl is an exported variable outside package refrence to all templates
//var Tmpl = template.Must(template.ParseGlob("../templates/*"))

var Tmpl = template.Must(template.ParseGlob(temlPath))

// Page is the struct used for each webpage
type Page struct {
	ID      int
	Title   string
	Content string
	Posts   []*Post
}

// Post is the struct used for each blog post
type Post struct {
	ID            int
	Title         string
	Content       string
	DatePublished time.Time
	Comments      []*Comment
}

// Comment is the struct used for each comment
type Comment struct {
	ID            int
	Author        string
	Comment       string
	DatePublished time.Time
}
