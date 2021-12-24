package main

import (
	"io"
	"os"
	"testing"
)

func getHandle(t *testing.T, s string) io.Reader {
	t.Helper()
	r, err := os.Open(s)
	if err != nil {
		t.Fatalf("failed to open testdata (%s): %v", s, err)
	}
	return r
}

func TestP1(t *testing.T) {
	tests := []struct {
		name    string
		data    io.Reader
		want    int64
		wantErr bool
	}{
		{
			"basic",
			getHandle(t, "testdata/input.txt"),
			150,
			false,
		}, {
			"aocd2",
			getHandle(t, "testdata/input2.txt"),
			0,
			false,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := p1(tc.data)
			if (err != nil) != tc.wantErr {
				t.Errorf("func %s goterr: %v wanted: %v", tc.name, err, tc.wantErr)
			}
			if got != tc.want {
				t.Errorf("func got: %d, want: %d", got, tc.want)
			}
		})
	}
}
