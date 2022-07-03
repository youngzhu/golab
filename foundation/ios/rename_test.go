package ios

import (
	"log"
	"os"
	"testing"
)

func TestRename(t *testing.T) {
	from, to := "renamefrom", "renameto"

	file, err := os.Create(from)
	if err != nil {
		t.Fatalf("open %q failed: %v", from, err)
	}
	if err = file.Close(); err != nil {
		t.Errorf("close %q failed: %v", from, err)
	}
	log.Println(file.Name())
	err = os.Rename(from, to)
	if err != nil {
		t.Fatalf("rename %q, %q failed: %v", to, from, err)
	}
	_, err = os.Stat(to)
	if err != nil {
		t.Errorf("stat %q failed: %v", to, err)
	}
}

func TestRenameOverwriteDest(t *testing.T) {
	from, to := "renamefrom", "renameto"

	toData := []byte("to")
	fromData := []byte("from")

	err := os.WriteFile(to, toData, 0777)
	if err != nil {
		t.Fatalf("write file %q failed: %v", to, err)
	}

	err = os.WriteFile(from, fromData, 0777)
	if err != nil {
		t.Fatalf("write file %q failed: %v", from, err)
	}
	err = os.Rename(from, to)
	if err != nil {
		t.Fatalf("rename %q, %q failed: %v", to, from, err)
	}

	_, err = os.Stat(from)
	if err == nil {
		t.Errorf("from file %q still exists", from)
	}
	if err != nil && !os.IsNotExist(err) {
		t.Fatalf("stat from: %v", err)
	}
	toFi, err := os.Stat(to)
	if err != nil {
		t.Fatalf("stat %q failed: %v", to, err)
	}
	if toFi.Size() != int64(len(fromData)) {
		t.Errorf(`"to" size = %d; want %d (old "from" size)`, toFi.Size(), len(fromData))
	}
}

func TestRenameFailed(t *testing.T) {
	from, to := "renamefrom", "renameto"

	err := os.Rename(from, to)
	switch err := err.(type) {
	case *os.LinkError:
		if err.Op != "rename" {
			t.Errorf("rename %q, %q: err.Op: want %q, got %q", from, to, "rename", err.Op)
		}
		if err.Old != from {
			t.Errorf("rename %q, %q: err.Old: want %q, got %q", from, to, from, err.Old)
		}
		if err.New != to {
			t.Errorf("rename %q, %q: err.New: want %q, got %q", from, to, to, err.New)
		}
	case nil:
		t.Errorf("rename %q, %q: expected error, got nil", from, to)
	default:
		t.Errorf("rename %q, %q: expected %T, got %T %v", from, to, new(os.LinkError), err, err)
	}
}

func TestRenameNotExisting(t *testing.T) {
	from, to := "doesnt-exist", "dest"

	os.Mkdir(to, 0777)

	if err := os.Rename(from, to); !os.IsNotExist(err) {
		t.Errorf("Rename(%q, %q) = %v; want an IsNotExist error", from, to, err)
	}
}

func TestRenameToDirFailed(t *testing.T) {
	from, to := "renamefrom", "renameto"

	os.Mkdir(from, 0777)
	os.Mkdir(to, 0777)

	err := os.Rename(from, to)
	switch err := err.(type) {
	case *os.LinkError:
		if err.Op != "rename" {
			t.Errorf("rename %q, %q: err.Op: want %q, got %q", from, to, "rename", err.Op)
		}
		if err.Old != from {
			t.Errorf("rename %q, %q: err.Old: want %q, got %q", from, to, from, err.Old)
		}
		if err.New != to {
			t.Errorf("rename %q, %q: err.New: want %q, got %q", from, to, to, err.New)
		}
	case nil:
		t.Errorf("rename %q, %q: expected error, got nil", from, to)
	default:
		t.Errorf("rename %q, %q: expected %T, got %T %v", from, to, new(os.LinkError), err, err)
	}
}

func TestRenameCaseDifference(pt *testing.T) {
	from, to := "renameFROM", "RENAMEfrom"
	tests := []struct {
		name   string
		create func() error
	}{
		{"dir", func() error {
			return os.Mkdir(from, 0777)
		}},
		{"file", func() error {
			fd, err := os.Create(from)
			if err != nil {
				return err
			}
			return fd.Close()
		}},
	}

	for _, test := range tests {
		pt.Run(test.name, func(t *testing.T) {
			if err := test.create(); err != nil {
				t.Fatalf("failed to create test file: %s", err)
			}

			if _, err := os.Stat(to); err != nil {
				// Sanity check that the underlying filesystem is not case sensitive.
				if os.IsNotExist(err) {
					t.Skipf("case sensitive filesystem")
				}
				t.Fatalf("stat %q, got: %q", to, err)
			}

			if err := os.Rename(from, to); err != nil {
				t.Fatalf("unexpected error when renaming from %q to %q: %s", from, to, err)
			}

			fd, err := os.Open(".")
			if err != nil {
				t.Fatalf("Open .: %s", err)
			}

			// Stat does not return the real case of the file (it returns what the called asked for)
			// So we have to use readdir to get the real name of the file.
			dirNames, err := fd.Readdirnames(-1)
			if err != nil {
				t.Fatalf("readdirnames: %s", err)
			}

			if dirNamesLen := len(dirNames); dirNamesLen != 1 {
				t.Fatalf("unexpected dirNames len, got %q, want %q", dirNamesLen, 1)
			}

			if dirNames[0] != to {
				t.Errorf("unexpected name, got %q, want %q", dirNames[0], to)
			}
		})
	}
}
