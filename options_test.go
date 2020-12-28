package wasataro_test

import (
	"testing"

	"github.com/onigra/wasataro"
)

func TestPortWithColon(t *testing.T) {
	// given
	o := &wasataro.Options{
		Response: "hello",
		Port:     3000,
	}

	// when
	expected := ":3000"
	actual := o.PortWithColon()

	// then
	if actual != expected {
		t.Errorf("actual: %s, expected: %s", actual, expected)
	}
}
