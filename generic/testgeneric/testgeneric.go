package testgeneric

import (
	"io"
	"os"
	"testing"
)

func GetHandle(t *testing.T, s string) io.Reader {
	t.Helper()
	r, err := os.Open(s)
	if err != nil {
		t.Fatalf("failed to open testdata (%s): %v", s, err)
	}
	return r
}
