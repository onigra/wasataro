package main

import (
	"flag"
	"fmt"
	"net/http"
	"strings"
)

// String Add ServeHTTP Function
type String string

func (s String) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method, r.URL)
	fmt.Fprint(w, s)
}

func main() {
	response := flag.String("response", "ok", "response")
	port := flag.String("port", "3000", "port")
	flag.Parse()

	http.Handle("/", String(*response))
	p := []string{":", *port}
	http.ListenAndServe(strings.Join(p, ""), nil)
}
