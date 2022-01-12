package main

import (
	"net/http"
	"os"
	"text/template"
	"yasir2000/go-web-dev-side-project-1/cms/chat"
)

var index = template.Must(template.ParseFiles("./index.html"))

func home(w http.ResponseWriter, r *http.Request) {
	index.Execute(w, nil)
}

func main() {
	go chat.DefaultHub.Start()

	http.HandleFunc("/", home)
	http.HandleFunc("/notify", chat.Notify)
	http.HandleFunc("/ws", chat.WSHandler)
	if os.Getenv("env") == "dev" {
		http.ListenAndServe(":3000", nil)
	} else {
		http.ListenAndServe(":80", nil)
	}

}
