package main

import (
	"net/http"
	"yasir2000/go-web-dev-side-project-1/cms/users/api"
)

func main() {
	http.HandleFunc("/", api.Doc)
	http.HandleFunc("/newpage", api.CreatePage)
	http.HandleFunc("/pages", api.AllPages)
	http.HandleFunc("/page/", api.GetPage)

	http.ListenAndServe(":3000", nil)
}
