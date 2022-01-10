package api

import (
	"encoding/json"
	"net/http"
	"strings"
	cms "yasir2000/go-web-dev-side-project-1/cms"
)

var pool = New()

// Doc lists all routes in this API
func Doc(w http.ResponseWriter, r *http.Request) {
	data := (map[string]string{
		"all_pages_url":   "/pages",
		"page_url":        "/pages/{id}",
		"create_page_url": "/newpage",
	})
	writeJSON(w, data)
}

//AllPages returns all the pages
func AllPages(w http.ResponseWriter, r *http.Request) {
	data, err := cms.GetPages()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, data)
}

// getPage gets a single page from this API
func GetPage(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimLeft(r.URL.Path, "/pages/")
	data, err := cms.GetPage(id)
	if err != nil {
		errJSON(w, err.Error(), http.StatusNotFound)
	}
	writeJSON(w, data)
}

// CreatePage creates a new post or page from Page construct
func CreatePage(w http.ResponseWriter, r *http.Request) {
	page := new(cms.Page)
	//data, err := ioutil.ReadAll(r.Body)
	err := json.NewDecoder(r.Body).Decode(page)
	if err != nil {
		errJSON(w, err.Error(), http.StatusInternalServerError)
	}

	// Convert data json to page. allocates a new value for it to point to
	//json.Unmarshal(data, page)
	id, err := cms.CreatePage(page)
	if err != nil {
		errJSON(w, err.Error(), http.StatusInternalServerError)
	}

	// return a JSON response from a map
	writeJSON(w, map[string]int{
		"user_id": id,
	})
}

// To marshal interface map to JSON
func writeJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	// turns data into byte[] of json
	//resJSON, err := json.MarshalIndent(data, "", "\t")

	// The below line will stream data into JSON stream to byte buffer then
	// to response writer
	//var b bytes.Buffer
	buf := pool.Get()
	defer pool.Put(buf)
	//err := json.NewEncoder(&b).Encode(data)
	err := json.NewEncoder(buf).Encode(data)
	if err != nil {
		errJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//w.Write(resJSON)
	buf.WriteTo(w)
}

// Create JSON error manually
func errJSON(w http.ResponseWriter, err string, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte("{\n\terror: " + err + "\n}\n"))
}
