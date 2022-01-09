package users

import (
	"encoding/json"
	"net/http"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	identityURL = "http://www.googleapis.com/oauth2/v2/userinfo"
	provider    = New()
	signingKey  = genRandBytes()
)

// New function creates a new oauth2 config
func New() *oauth2.Config {
	return &oauth2.Config{
		ClientID:     os.Getenv("GOOGLE_KEY"),
		ClientSecret: os.Getenv("GOOGLE_SECRET"),
		Endpoint:     google.Endpoint,
		RedirectURL:  "http://localhost:3000/auth/gplus/callback",
		Scopes:       []string{"email", "profile"},
	}
}

// AuthURLHandler just an http handler which redirects
// the user to correct Oauth sign in page for provider (google)

func AuthURLHandler(w http.ResponseWriter, r *http.Request) {
	authURL := provider.AuthCodeURL("", oauth2.AccessTypeOffline)
	http.Redirect(w, r, authURL, http.StatusTemporaryRedirect)
}

// CallbackURLHandler is an http handler that handles all of the Oauth coming back flow
func CallbackURLHandler(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	token, err := provider.Exchange(oauth2.NoContext, code)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return

	}
	client := provider.Client(oauth2.NoContext, token)
	resp, err := client.Get(identityURL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	defer resp.Body.Close()

	user := make(map[string]string)
	// This decode will error, since it can't decode every value returned
	// from the API. However, we don't need to worry about this, since it can
	// correctly decode our user's email address.
	json.NewDecoder(resp.Body).Decode(&user)

	email := user["email"]
	genToken(w, email)
}

//genToken is and http responder that will take user as string and returns
// http response of the generated token by jwt protocol HS256
func genToken(w http.ResponseWriter, user string) {
	token := jwt.New(jwt.SigningMethodES256)
	tokenClaims := token.Claims.(jwt.MapClaims)
	// Who is the user this token for
	tokenClaims["sub"] = user
	// Expiration of this token
	tokenClaims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	// When this token was issued
	tokenClaims["iat"] = time.Now().Unix()

	tokenString, err := token.SignedString(signingKey)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// World's laziest way to issue a JSON response
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{\n\ttoken: " + tokenString + "\n"))
}

// VerifyToken gets the token from an HTTP request, and ensures that it's
// valid. It'll return the user's username as a string
// OAuth2Extractor extract and parse a JWT token from an HTTP request.
// This behaves the same as Parse, but accepts a request and an extractor
// instead of a token string.  The Extractor interface allows you to define
// the logic for extracting a token.  Several useful implementations are provided.
//
// You can provide options to modify parsing behavior
func VerifyToken(r *http.Request) (string, error) {
	// request is a sub-package of jwt-go
	token, err := request.ParseFromRequest(r, request.OAuth2Extractor, func(token *jwt.Token) (interface{}, error) {

		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return signingKey, nil
	})

	if err != nil {
		return "", err
	}
	if !token.Valid {
		return "", jwt.ErrInvalidKey
	}

	tokenClaims := token.Claims.(jwt.MapClaims)
	return tokenClaims["sub"].(string), nil
}
