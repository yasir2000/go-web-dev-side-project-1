package chat

import (
	"encoding/json"

	"net/http"
)

// Notify is an http handler which sends
// notifications to the client based on the
// contents of the POST request.
func Notify(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		if r.Header.Get("Content-Type") == "application/json" {
			jq := make(map[string]string)
			json.NewDecoder(r.Body).Decode(&jq)
			msg := jq["message"]

			if msg == "" {
				http.Error(w, `Messages must have the key "messages"`, http.StatusBadRequest)
				return
			}
			DefaultHub.Echo <- msg
			w.Write([]byte("Sent message successfully"))
			return
		}

		msg := r.FormValue("message")
		r.ParseForm()

		if msg == "" {
			http.Error(w, "No message found in request", http.StatusBadRequest)
			return
		}

		DefaultHub.Echo <- msg
		w.Write([]byte("Sent message successfully"))
		return
	default:
		http.Error(w, "Only POST method supported", http.StatusMethodNotAllowed)
		return
	}
}
