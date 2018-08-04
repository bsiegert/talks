// +build !appengine

package main

import (
	"flag"
	"fmt"
	"net/http"
)

var addr *string = flag.String("addr", ":8080", "host:port to listen on")

func main() {
	flag.Parse()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello World!")
	})

	http.ListenAndServe(*addr, nil)
}
