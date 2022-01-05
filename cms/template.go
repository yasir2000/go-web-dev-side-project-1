package cms

import (
	"html/template"
)

//Tmpl is an exported variable outside package
var Tmpl = template.Must(template.ParseGlob("../templates/*"))

type Page struct {
	Title   string
	Content string
}
