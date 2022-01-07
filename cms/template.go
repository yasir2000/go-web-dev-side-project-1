package cms

import (
	"html/template"
	"time"
)

//Tmpl is an exported variable outside package refrence to all templates
var Tmpl = template.Must(template.ParseGlob("../templates/*"))

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
