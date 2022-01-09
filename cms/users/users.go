package users

import (
	"errors"
	"time"

	"github.com/boltdb/bolt"
	"golang.org/x/crypto/bcrypt"
)

var (
	// DB is the reference to DB, which contains user data,
	DB = newDB()

	// ErrUserAlreadyExists is the error thrown when a user
	// attempts to create a new user in the DB with a duplicate username.
	ErrUserAlreadyExists = errors.New("users: username already exists")

	// ErrUserNotFound is an error thrown when a user can't be found in the store DB
	ErrUserNotFound = errors.New("users: user not found")
)

// Store is a very simple in memory DB, it will b eused to store
// users. Protected by read-wrote mutex (RWMutex), so that two goroutines can't
// modify the underlying map at same time (since maps are not safe concurrent
// in Go)
//
// -------------------------------------------
// Store is a reference to our BoltDB instance that contains two seperate
// internal stores: a user store, and a session store.

type Store struct {
	//rwm *sync.RWMutex
	//m   map[string]string
	DB       *bolt.DB
	Users    string
	Sessions string
}

//newDB is a convenience method to initialize our in memory DB when
// program starts

func newDB() *Store {
	// Create or open the DB
	db, err := bolt.Open("users.db", 0600, &bolt.Options{
		Timeout: 1 * time.Second,
	})
	if err != nil {
		panic(err)
	}
	//return &Store{
	//	rwm: &sync.RWMutex{},
	//	m:   make(map[string]string),

	//  -----------------
	// Create the Users bucket
	db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("Users"))
		if err != nil {
			return err
		}
		return nil
	})

	// Create the Sessions bucket
	db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("Sessions"))
		if err != nil {
			return err
		}
		return nil
	})

	return &Store{
		DB:       db,
		Users:    "Users",
		Sessions: "Sessions",
	}

}

//NewUser functions accepts a username and password and creates a new
// user in DB
func NewUser(username string, password string) error {
	err := exists(username)
	if err != nil {
		return err
	}

	//DB.rwm.Lock()
	//defer DB.rwm.Unlock()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	//DB.m[username] = string(hashedPassword)
	//return nil

	// ------------

	return DB.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(DB.Users))
		return b.Put([]byte(username), hashedPassword)
	})
}

// AuthenticateUser accepts a username and password , then checks
// if given passwored matches the hashed password. It returns nil on success
// and an error on failure
func AuthenticateUser(username string, password string) error {
	var hashedPassword []byte
	DB.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(DB.Users))
		hashedPassword = b.Get([]byte(username))
		return nil
	})
	//DB.rwm.RLock()
	//defer DB.rwm.RUnlock()
	// -----------
	if hashedPassword == nil {
		return ErrUserNotFound
	}
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	//hashedPassword := DB.m[username]
	//err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	//return err
}

//OverrideOldPassword overrides the old password with new one
// will be used when resetting password.
func OverrideOldPassword(username string, password string) error {

	//DB.rwm.Lock()
	//defer DB.rwm.Unlock()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	return DB.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(DB.Users))
		return b.Put([]byte(username), hashedPassword)
	})

	//DB.m[username] = string(hashedPassword)
	//return nil
}

// exists is an internal utility function for ensuring the usernames are
// unique
func exists(username string) error {
	var result []byte
	DB.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(DB.Users))
		result = b.Get([]byte(username))
		return nil
	})
	if result != nil {
		return ErrUserAlreadyExists
	}

	return nil
}
