package wasataro

import (
	"fmt"
	"os"
)

func usageText() string {
	return `Usage: wasataro [<flags>]

Flags:
  -response (Optional) Modify response body. default: ok
  -port     (Optional) Change listen port.   default: 3000`
}

// Usage print usage message
func Usage() func() {
	return func() {
		fmt.Fprintf(os.Stderr, "%s\n", usageText())
	}
}
