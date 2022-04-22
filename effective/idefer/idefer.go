package idefer

import (
	"io"
	"os"
)

// LIFO，如果有多个defer，后来的先执行

// Contents returns the file's contents as a string
func Contents(filename string) (string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer f.Close() // 在函数结束前执行

	var result []byte
	buf := make([]byte, 100)
	for {
		n, err := f.Read(buf[0:])
		result = append(result, buf[0:n]...)
		if err != nil {
			if err == io.EOF {
				break
			}
			return "", err
		}
	}

	return string(result), nil
}
