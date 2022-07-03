package create_temp

import (
	"log"
	"os"
	"testing"
)

func TestCreateTemp(t *testing.T) {
	f, err := os.CreateTemp("./", "example")
	if err != nil {
		log.Fatal(err)
	}
	//defer os.Remove(f.Name()) // clean up

	if _, err := f.Write([]byte("content")); err != nil {
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}

func TestCreateTemp_suffix(t *testing.T) {
	/*
		tests := []struct{ pattern, prefix, suffix string }{
				{"tempfile_test", "tempfile_test", ""},
				{"tempfile_test*", "tempfile_test", ""},
				{"tempfile_test*xyz", "tempfile_test", "xyz"},
			}
	*/
	f, err := os.CreateTemp("./", "example.*.txt")
	if err != nil {
		log.Fatal(err)
	}
	//defer os.Remove(f.Name()) // clean up

	if _, err := f.Write([]byte("content")); err != nil {
		f.Close()
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}
