package users

import (
	"crypto/rand"
	"encoding/base64"
	"html/template"
	"net/http"
	"net/smtp"
	"os"
)

// validTokens is an array of the currently valid tokens
var validTokens []string

const passwordForm = `
<h1>Enter your new password</h1>
<form action="/reset" method="POST">
	<input type="hidden" name="email" value="{{ . }}" required>

	<label for="password">Password</label>
	<input type="password" name="password" required>

	<input type="submit" value="Submit">
</form>
`

// genRandBytes generates 32 byte long string of random bytes
func genRandBytes() []byte {
	b := make([]byte, 24)
	_, err := rand.Read(b)
	if err != nil {
		//Panic if we can't
	}
	return []byte(base64.URLEncoding.EncodeToString(b))
}

//SendPasswordresetEmail sends a password reset email to a given user
func SendPasswordResetEmail(email string) error {
	token := string(genRandBytes())
	validTokens = append(validTokens, token)
	resetLink := "http://localhost:3000/reset?user=" + email + "&token" + token

	username := os.Getenv("GMAIL_USERNAME")
	password := os.Getenv("GMAIL_PASSWORD")
	auth := smtp.PlainAuth("smtp.gmail.com:587", username, password, "smtp.gmail.com")

	return smtp.SendMail("smtp.gmail.com:587", auth, username, []string{email}, []byte("Click here to reset your password:"+resetLink))
}

//ResetPassword is the HTTP handler that handles the password reset flow

func ResetPassword(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		var isValid bool
		var index int

		//Get the query string parameters
		email := r.URL.Query().Get("user")
		token := r.URL.Query().Get("token")

		// Ensure the token is valid one
		for i, ok := range validTokens {
			if ok == token {
				isValid = true
				index = i
			}

		}

		if isValid != true {
			http.Error(w, "Token is not valid", http.StatusUnauthorized)
			return

		}
		//Delete the token
		validTokens = append(validTokens[:index], "")

		//Render the reset template
		t, _ := template.New("password").Parse(passwordForm)
		t.Execute(w, email)
		return
	case "POST":
		//Get the new password from a form value
		password := r.FormValue("password")
		email := r.FormValue("email")
		r.ParseForm()

		// Reset the password by overwriting the old password
		err := OverrideOldPassword(email, password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write([]byte("Successfully reset password"))

	}
}
