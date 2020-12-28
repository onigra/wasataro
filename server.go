package wasataro

import (
	"fmt"
	"net/http"
)

// String Add ServeHTTP Function
type String string

func (s String) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method, r.URL)
	fmt.Fprint(w, s)
}

// Start start http server
func Start(o Options) {
	http.Handle("/", String(o.Response))
	http.ListenAndServe(o.PortWithColon(), nil)
}
