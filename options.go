package wasataro

import (
	"strconv"
	"strings"
)

// Options cli option
type Options struct {
	Response string
	Port     int
}

// PortWithColon return port with colon
// 3000 -> ":3000"
func (o Options) PortWithColon() string {
	p := []string{":", strconv.Itoa(o.Port)}
	return strings.Join(p, "")
}
