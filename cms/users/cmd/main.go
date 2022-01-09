package main

import (
	"fmt"
	"net/http"
	"text/template"
	"yasir2000/go-web-dev-side-project-1/cms/users"
)

const loginTemplate = `
<h1>Enter your username and password</h1>
<form action="/" method="POST">
	<input type="text" name="user" required>

	<label for="password">Password</label>
	<input type="password" name="password" required>

	<input type="submit" value="Submit">
</form>
`

func authHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		t, _ := template.New("login").Parse(loginTemplate)
		t.Execute(w, nil)
	case "POST":
		user := r.FormValue("user")
		pass := r.FormValue("password")
		err := users.AuthenticateUser(user, pass)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		users.SetSession(w, user)
		w.Write([]byte("Signed in successfully"))
	}

}

func restrictedHandler(w http.ResponseWriter, r *http.Request) {
	user := users.GetSession(w, r)
	w.Write([]byte(user))
}

func main() {

	username, password := "akarama1933", "querty123"

	//username, password := os.Getenv("GMAIL_USERNAME"), "qwerty123"

	err := users.NewUser(username, password)
	if err != nil {
		fmt.Printf("Couldn't create user: %s\n", err.Error())
		return
	}

	// err = users.AuthenticateUser(username, password)
	// if err != nil {
	// 	fmt.Printf("Couldn't authenticate user : %s\n", err.Error())
	// 	return
	// }

	//	fmt.Println("Successfuly created and authenticated user %s", username)
	fmt.Printf("Succesfully created and authenticated user \033[32m%s\033[0m\n", username)

	// Send reset email
	// err = users.SendPasswordResetEmail(username)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	//http.HandleFunc("/reset", users.ResetPassword)
	http.HandleFunc("/", authHandler)
	http.HandleFunc("/restricted", restrictedHandler)
	http.ListenAndServe(":3000", nil)
}
