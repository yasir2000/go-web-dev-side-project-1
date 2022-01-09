package users

import (
	"errors"
	"sync"

	"golang.org/x/crypto/bcrypt"
)

var (
	// DB is the reference to DB, which contains user data,
	DB = newDB()

	// ErrUserAlreadyExists is the error thrown when a user
	// attempts to create a new user in the DB with a duplicate username.
	ErrUserAlreadyExists = errors.New("users: username already exists")
)

// Store is a very simple in memory DB, it will b eused to store
// users. Protected by read-wrote mutex (RWMutex), so that two goroutines can't
// modify the underlying map at same time (since maps are not safe concurrent
// in Go)

type Store struct {
	rwm *sync.RWMutex
	m   map[string]string
}

//newDB is a convenience method to initialize our in memory DB when
// program starts

func newDB() *Store {
	return &Store{
		rwm: &sync.RWMutex{},
		m:   make(map[string]string),
	}
}

//NewUser functions accepts a username and password and creates a new
// user in DB
func NewUser(username string, password string) error {
	err := exists(username)
	if err != nil {
		return err
	}

	DB.rwm.Lock()
	defer DB.rwm.Unlock()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	DB.m[username] = string(hashedPassword)
	return nil
}

// AuthenticateUser accepts a username and password , then checks
// if given passwored matches the hashed password. It returns nil on success
// and an error on failure
func AuthenticateUser(username string, password string) error {

	DB.rwm.RLock()
	defer DB.rwm.RUnlock()

	hashedPassword := DB.m[username]
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	return err
}

//OverrideOldPassword overrides the old password with new one
// will be used when resetting password.
func OverrideOldPassword(username string, password string) error {
	DB.rwm.Lock()
	defer DB.rwm.Unlock()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	DB.m[username] = string(hashedPassword)
	return nil
}

// exists is an internal utility function for ensuring the usernames are
// unique
func exists(username string) error {
	DB.rwm.RLock()
	defer DB.rwm.RUnlock()

	if DB.m[username] != "" {
		return ErrUserAlreadyExists
	}
	return nil
}
