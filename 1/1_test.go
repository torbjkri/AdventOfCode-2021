package dec1

import (
	//"regexp"
	"os"
	"testing"
)

// Test empty filename
func TestReadFileEmptyFilename(t *testing.T) {
	filename := ""
	data, err := ReadFile(filename)
	if data != nil || err == nil {
		t.Fatalf(`ReadFile("") = %q, %v, want ""`, data, err)
	}
}

func TestReadFileInvalidFilename(t *testing.T) {
	filename := "file"
	data, err := ReadFile(filename)
	if data != nil || err == os.fs.PathError {
		t.Fatalf(`ReadFile("file") = %q, %v, want ""`, data, err)
	}
}
