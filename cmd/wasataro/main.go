package main

import (
	"flag"

	"github.com/onigra/wasataro"
)

func main() {
	flag.Usage = wasataro.Usage()
	response := flag.String("response", "ok", "response")
	port := flag.Int("port", 3000, "port")
	flag.Parse()

	options := &wasataro.Options{
		Response: *response,
		Port:     *port,
	}

	wasataro.Start(*options)
}
