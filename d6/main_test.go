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
		daynum  int
		want    int64
		wantErr bool
	}{
		{
			"basic",
			testgeneric.GetHandle(t, "testdata/input1.txt"),
			80,
			5934,
			false,
		},
		{
			"aocd6",
			testgeneric.GetHandle(t, "testdata/input2.txt"),
			80,
			0,
			false,
		},
		{
			"aocd6p2",
			testgeneric.GetHandle(t, "testdata/input2.txt"),
			256,
			0,
			false,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := p1(tc.data, tc.daynum)
			if (err != nil) != tc.wantErr {
				t.Errorf("func %s goterr: %v wanted: %v", tc.name, err, tc.wantErr)
			}
			if got != tc.want {
				t.Errorf("func got: %d, want: %d", got, tc.want)
			}
		})
	}
}
