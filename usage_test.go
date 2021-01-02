package wasataro_test

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/onigra/wasataro"
)

func TestUsage(t *testing.T) {
	// Setup
	// catch stderr
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}

	stderr := os.Stderr
	os.Stderr = w

	// And
	expected := `Usage: wasataro [<flags>]

Flags:
  -response (Optional) Modify response body. default: ok
  -port     (Optional) Change listen port.   default: 3000`

	// when
	f := wasataro.Usage()
	f()

	// Teardown
	os.Stderr = stderr
	w.Close()

	var buf bytes.Buffer
	io.Copy(&buf, r)

	// then
	raw := buf.String()
	actual := strings.TrimRight(raw, "\n")

	if actual != expected {
		t.Errorf("actual: %s, expected: %s", actual, expected)
	}
}
