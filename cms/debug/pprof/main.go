// This examples shows the data output by pprof
//
// To see all the URLs available to you when you run pprof, visit
// http://localhost:3000/debug/pprof.
//
// To have some numbers to look at, apply some load to your server:
//
//   wrk -d10 -c20 -t10 "http://localhost:3000/"
//
package main

import (
	"net/http"
	_ "net/http/pprof"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, world!"))
	})

	http.ListenAndServe(":3000", nil)
}

// docker run --rm -v `pwd`:/data williamyeh/wrk -c 100 http://192.168.1.46:3000
// docker run --rm -v `pwd`:/data williamyeh/wrk -t12 -c400 -d30s --latency http://192.168.1.46:3000
// go tool pprof http://localhost:3000/debug/pprof/profile
