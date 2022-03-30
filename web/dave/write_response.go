package dave

import (
	"fmt"
	"io"
)

// 去掉对错误的重复处理

type Header struct {
	Key, Value string
}

type Status struct {
	Code   int
	Reason string
}

// 这里对err的判断就是重复操作

func WriteResponse(w io.Writer, st Status, headers []Header, body io.Reader) error {
	_, err := fmt.Fprintf(w, "HTTP/1.1 %d %s\r\n", st.Code, st.Reason)
	if err != nil {
		return err
	}

	for _, h := range headers {
		_, err := fmt.Fprintf(w, "%s: %s\r\n", h.Key, h.Value)
		if err != nil {
			return err
		}
	}

	if _, err := fmt.Fprint(w, "\r\n"); err != nil {
		return err
	}

	_, err = io.Copy(w, body)
	return err
}

/******************************************/
/*************  对err优化处理  **************/
/******************************************/

type errWriter struct {
	io.Writer
	err error
}

func (e errWriter) Write(buf []byte) (int, error) {
	// 优化的处理在这里
	if e.err != nil {
		return 0, e.err
	}
	var n int
	n, e.err = e.Write(buf)
	return n, nil
}

func WriteResponseX(w io.Writer, st Status, headers []Header, body io.Reader) error {
	ew := &errWriter{Writer: w}
	fmt.Fprintf(ew, "HTTP/1.1 %d %d\r\n", st.Code, st.Reason)

	for _, h := range headers {
		fmt.Fprintf(ew, "%s: %s\r\n", h.Key, h.Value)
	}

	fmt.Fprintf(ew, "\r\n")
	io.Copy(ew, body)
	return ew.err
}
