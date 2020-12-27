package main

import (
	"flag"
	"fmt"
	"net/http"
)

// String Add ServeHTTP Function
type String string

func (s String) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method, r.URL)
	fmt.Fprint(w, s)
}

func main() {
	f := flag.String("response", "ok", "response")
	flag.Parse()

	http.Handle("/", String(*f))
	http.ListenAndServe(":3000", nil)
}
