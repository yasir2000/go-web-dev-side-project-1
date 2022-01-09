package main

import (
	"fmt"
	"net/http"
	"yasir2000/go-web-dev-side-project-1/cms/users"
)

func main() {

	username, password := "akarama1932", "Ytds2012"

	//username, password := os.Getenv("GMAIL_USERNAME"), "qwerty123"

	err := users.NewUser(username, password)
	if err != nil {
		fmt.Printf("Couldn't create user: %s\n", err.Error())
		return
	}

	err = users.AuthenticateUser(username, password)
	if err != nil {
		fmt.Printf("Couldn't authenticate user : %s\n", err.Error())
		return
	}

	fmt.Println("Successfuly created and authenticated user %s", username)

	// Send reset email
	err = users.SendPasswordResetEmail(username)
	if err != nil {
		fmt.Println(err)
	}

	http.HandleFunc("/reset", users.ResetPassword)
	http.ListenAndServe(":3000", nil)
}
