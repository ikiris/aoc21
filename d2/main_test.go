package main

import (
	"io"
	"testing"

	"github.com/ikiris/aoc21/generic/testgeneric"
)

func TestP1(t *testing.T) {
	tests := []struct {
		name    string
		data    io.Reader
		want    int64
		wantErr bool
	}{
		{
			"basic",
			testgeneric.GetHandle(t, "testdata/input.txt"),
			150,
			false,
		}, {
			"aocd2",
			testgeneric.GetHandle(t, "testdata/input2.txt"),
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
func TestP2(t *testing.T) {
	tests := []struct {
		name    string
		data    io.Reader
		want    int64
		wantErr bool
	}{
		{
			"basic",
			testgeneric.GetHandle(t, "testdata/input.txt"),
			900,
			false,
		}, {
			"aocd2",
			testgeneric.GetHandle(t, "testdata/input2.txt"),
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
