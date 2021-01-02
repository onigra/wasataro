package main

import (
	"flag"

	"github.com/onigra/wasataro"
)

func main() {
	response := flag.String("response", "ok", "(Optional) Modify response body.")
	port := flag.Int("port", 3000, "(Optional) Change listen port.")
	flag.Parse()

	options := &wasataro.Options{
		Response: *response,
		Port:     *port,
	}

	wasataro.Start(*options)
}
