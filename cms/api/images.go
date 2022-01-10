package api

import (
	"io"
	"net/http"
	"os"
	"strings"
)

var here = os.Getenv("GOPATH") + "\\src\\github.com\\yasir2000\\go-web-dev-side-project-1\\cms\\api\\cmd\\images"

// UploadImage allows to upload an image as a request
func UploadImage(w http.ResponseWriter, r *http.Request) {
	file, header, err := r.FormFile("image-")
	if err != nil {
		errJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// wait till uploading file finishes then closes file
	defer file.Close()

	// Creates an os file, it is created with mode 0666
	out, err := os.Create(here + header.Filename)
	if err != nil {
		errJSON(w, err.Error(), http.StatusInternalServerError)
	}

	// waits till file is created the closes file
	defer out.Close()

	// Copies file, from src to dst
	_, err = io.Copy(out, file)
	if err != nil {
		errJSON(w, err.Error(), http.StatusInternalServerError)
	}

	// Creates a JSON http response from map of strings include filename
	writeJSON(w, map[string]string{
		"filename": header.Filename,
	})
}

// SHowImage shows/displays the image based on filename provided from path
func ShowImage(w http.ResponseWriter, r *http.Request) {
	name := strings.TrimLeft(r.URL.Path, "/image/")
	file, err := os.Open(here + name)
	if err != nil {
		errJSON(w, err.Error(), http.StatusNotFound)
		return
	}

	buf := pool.Get()
	defer pool.Put(buf)
	_, err = io.Copy(buf, file)
	if err != nil {
		errJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "image/jpeg")
	buf.WriteTo(w)
}
