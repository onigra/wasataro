package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/onigra/wasataro"
)

// String Add ServeHTTP Function
type String string

func (s String) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method, r.URL)
	fmt.Fprint(w, s)
}

func main() {
	response := flag.String("response", "ok", "response")
	port := flag.Int("port", 3000, "port")
	flag.Parse()

	options := &wasataro.Options{
		Response: *response,
		Port:     *port,
	}

	http.Handle("/", String(options.Response))
	http.ListenAndServe(options.PortWithColon(), nil)
}
