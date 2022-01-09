package main

import (
	"fmt"
	"log"
	"net/http"
	"yasir2000/go-web-dev-side-project-1/cms/middleware"

	"golang.org/x/net/context"
	//"goa.design/goa/http/middleware"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Executing ...")
	w.Write([]byte("Hello"))
}

func tricky() string {
	defer log.Println("String 2")
	return "String 1" // will print 2 1
}

func lastInFirstOut() {
	for i := 0; i < 4; i++ {
		defer log.Println(i)
	} // Will print 3 2 1 0
}

func panicker(w http.ResponseWriter, r *http.Request) {
	panic(middleware.ErrInvalidEmail)
}

// this fix replaces a bar with type byte (bar) as bar is ctx of a map "vals" type
func withContext(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	bar := ctx.Value("vals").(middleware.Values).Get("foo")
	w.Write([]byte(bar))

}

func main() {
	// sum := middleware.Add(1, 2, 3)
	// fmt.Println(sum)

	// chain := &middleware.Chain{0}
	// sum2 := chain.AddNext(1).AddNext(2).AddNext(3).Finally(0)
	// fmt.Println(sum2)

	// http.Handle("/", middleware.Next(hello))
	// http.ListenAndServe(":3000", nil)
	logger := middleware.CreateLogger("logger-file")
	http.Handle("/", middleware.Time(logger, hello))
	http.Handle("/panic", middleware.Recover(panicker))
	http.Handle("/context", middleware.PassContext(withContext))
	//http.Handle("/panic", middleware.Time(logger, hello))
	http.ListenAndServe(":3000", nil)
}
