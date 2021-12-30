package main

import (
	"io"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
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
			getHandle(t, "testdata/input1.txt"),
			198,
			false,
		}, {
			"aocd3",
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

func TestZipfilter(t *testing.T) {
	tests := []struct {
		name       string
		intList    []int64
		pos        int
		comparable func(a, b int) bool
		want       []int64
		wantErr    bool
	}{
		{
			"basic",
			[]int64{4, 30, 22, 23, 21, 15, 7, 28, 16, 25, 2, 10},
			4,
			func(a, b int) bool {
				return a >= b
			},
			[]int64{23},
			false,
		}, {
			"basic-2",
			[]int64{4, 30, 22, 23, 21, 15, 7, 28, 16, 25, 2, 10},
			4,
			func(a, b int) bool {
				return a < b
			},
			[]int64{10},
			false,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := zipfilter(tc.intList, tc.pos, tc.comparable)
			if (err != nil) != tc.wantErr {
				t.Errorf("func %s goterr: %v wanted: %v", tc.name, err, tc.wantErr)
			}
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("func diff(-got,+want):%v", diff)
			}
		})
	}
}

func TestP2(t *testing.T) {
	tests := []struct {
		name    string
		data    io.Reader
		want    int64
		wantErr bool
	}{
		{
			"basic",
			getHandle(t, "testdata/input1.txt"),
			230,
			false,
		},
		{
			"aocd3",
			getHandle(t, "testdata/input2.txt"),
			0,
			false,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := p2(tc.data)
			if (err != nil) != tc.wantErr {
				t.Errorf("func %s goterr: %v wanted: %v", tc.name, err, tc.wantErr)
			}
			if got != tc.want {
				t.Errorf("func got: %d, want: %d", got, tc.want)
			}
		})
	}
}
