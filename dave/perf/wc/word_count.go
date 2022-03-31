package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"unicode"
)

func readByte(r io.Reader) (rune, error) {
	var buf [1]byte
	_, err := r.Read(buf[:])
	return rune(buf[0]), err
}

func main() {
	fileName := os.Args[1]
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("could not open file %q: %v", fileName, err)
	}

	words := 0
	inword := false
	b := bufio.NewReader(f)
	for {
		r, err := readByte(b)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("could not read file %q: %v", fileName, err)
		}
		if unicode.IsSpace(r) && inword {
			words++
			inword = false
		}
		inword = unicode.IsLetter(r)
	}
	fmt.Printf("%q: %d words\n", fileName, words)
}
