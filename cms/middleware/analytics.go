package middleware

import (
	"errors"
	"log"
	"net/http"
	"os"
	"time"

	"golang.org/x/net/context"
)

var (
	ErrInvalidID    = errors.New("Invalid ID")
	ErrInvalidEmail = errors.New("Invalid email")
)

// //Add is a variadic function that adds up numbers
// func Add(nums ...int) int {
// 	sum := 0
// 	for _, num := range nums {
// 		sum += num
// 	}
// 	return sum
// }

// // Chain holds our sum
// type Chain struct {
// 	Sum int
// }

// // AddNext is  chainable sum function
// func (c *Chain) AddNext(num int) *Chain {
// 	c.Sum += num
// 	return c
// }

// // Finally is for use at the end of the chain, and returns the sum
// func (c *Chain) Finally(num int) int {
// 	return c.Sum + num
// }

// CreateLogger creates a new logger that writes to the given filename.
func CreateLogger(filename string) *log.Logger {
	file, err := os.OpenFile(filename+".log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)

	}
	logger := log.New(file, "", log.Ldate|log.Ltime|log.Lshortfile)
	return logger
}

// Recover recovers from any panicking goroutine
func Recover(next http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			err := recover()
			if err != nil {
				switch err {
				case ErrInvalidEmail:
					http.Error(w, ErrInvalidEmail.Error(), http.StatusUnauthorized)
				case ErrInvalidID:
					http.Error(w, ErrInvalidID.Error(), http.StatusUnauthorized)
				default:
					http.Error(w, "Unknown error, recovered from panic", http.StatusInternalServerError)
				}
			}
		}()
		next.ServeHTTP(w, r)
	})
}

// Next runs the next function in the chain, HandleFunc is a type here to be returned

// Time runs the next function in the chain
func Time(logger *log.Logger, next http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//fmt.Println("Before")
		start := time.Now()
		//fmt.Println(start)
		next.ServeHTTP(w, r)
		elapsed := time.Since(start)
		//fmt.Println("After")
		//fmt.Println(elapsed)
		logger.Println(elapsed)
	})
}

type Values struct {
	m map[string]string
}

func (v Values) Get(key string) string {
	return v.m[key]
}

// PassContext is used to pass values between middleware.
type PassContext func(ctx context.Context, w http.ResponseWriter, r *http.Request)

// ServeHTTP satisfies the http.Handler interface.
func (fn PassContext) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	v := Values{map[string]string{
		"foo": "bar",
	}}
	ctx := context.WithValue(context.Background(), "vals", v)
	fn(ctx, w, r)
}
