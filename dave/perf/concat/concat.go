package concat

import (
	"bytes"
	"fmt"
	"strings"
	"time"
)

//const (
var (
	ID   = "youngzy"
	Addr = "127.0.0.1"
)

func Plus() string {
	s := ID
	s += " " + Addr
	s += " " + time.Now().String()
	return s
}

func Buffer() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "%s %s %v", ID, Addr, time.Now())
	return b.String()
}

func Sprintf() string {
	return fmt.Sprintf("%s %s %v", ID, Addr, time.Now())
}

func Bytes() string {
	b := make([]byte, 0, 40)
	b = append(b, ID...)
	b = append(b, ' ')
	b = append(b, Addr...)
	b = append(b, ' ')
	b = time.Now().AppendFormat(b, "2006-01-02 15:04:05.999999999 -0700 MST")
	return string(b)
}
func Builder() string {
	var b strings.Builder
	b.WriteString(ID)
	b.WriteString(" ")
	b.WriteString(Addr)
	b.WriteString(" ")
	b.WriteString(time.Now().String())
	return b.String()
}
