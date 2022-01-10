package main

import (
	"net/http"
	"os"
	"yasir2000/go-web-dev-side-project-1/cms/api"
)

func main() {
	os.Mkdir("images", 0777)

	http.HandleFunc("/", api.Doc)
	http.HandleFunc("/image/", api.ShowImage)
	http.HandleFunc("/newpage", api.CreatePage)
	http.HandleFunc("/pages", api.AllPages)
	http.HandleFunc("/page/", api.GetPage)
	http.HandleFunc("/upload", api.UploadImage)

	http.ListenAndServe(":3000", nil)

}
