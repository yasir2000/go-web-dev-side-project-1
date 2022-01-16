package cms

import (
	"net/http"
	"strings"
	"time"
)

// ServeIndex function serves as handler of the index page.
func ServeIndex(w http.ResponseWriter, r *http.Request) {
	p := &Page{
		Title:   "Go Projects CMS",
		Content: "Welcome to our home page!",
		Posts: []*Post{
			&Post{
				Title:         "Hello, World",
				Content:       "Hello world! Thanks for comming to this site.",
				DatePublished: time.Now(),
			},
			&Post{
				Title:         "A Post with Comments.",
				Content:       "Here's a conversational post. It could attract comments.",
				DatePublished: time.Now().Add(-time.Hour),
				Comments: []*Comment{
					&Comment{
						Author:        "Ben Transfer",
						Comment:       "Nevermind, I guess I just commented on my post ...",
						DatePublished: time.Now().Add(-time.Hour / 2),
					},
				},
			},
		},
	}
	Tmpl.ExecuteTemplate(w, "page", p)
}

//HandleNew function handles form logic
func HandleNew(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		Tmpl.ExecuteTemplate(w, "new", nil)
	case "POST":
		title := r.FormValue("title")
		content := r.FormValue("content")
		contentType := r.FormValue("content-type")
		r.ParseForm()

		if contentType == "page" {
			p := &Page{
				Title:   title,
				Content: content,
			}

			_, err := CreatePage(p)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			Tmpl.ExecuteTemplate(w, "page", p)
			return
		}

		if contentType == "post" {
			Tmpl.ExecuteTemplate(w, "post", &Post{
				Title:   title,
				Content: content,
			})
			return
		}
	default:
		http.Error(w, "Method not found: "+r.Method, http.StatusMethodNotAllowed)
	}
}

// ServePage serves as handler for page based on route matched. This will match
// a URL beginning with /page
func ServePage(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimLeft(r.URL.Path, "/page/")

	if path == "" {
		pages, err := GetPages()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		Tmpl.ExecuteTemplate(w, "pages", pages)
		return

	}
	//Since we've already got this route from our mock data, we will re-use it.
	page, err := GetPage(path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	Tmpl.ExecuteTemplate(w, "page", page)
}
