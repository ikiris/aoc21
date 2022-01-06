package main

import (
	"io"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/ikiris/aoc21/generic/testgeneric"
)

func TestP1(t *testing.T) {
	tests := []struct {
		name    string
		data    io.Reader
		want    int
		wantErr bool
	}{
		{
			"basic",
			testgeneric.GetHandle(t, "testdata/input1.txt"),
			26,
			false,
		},
		{
			"aocd8",
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

func TestDemap(t *testing.T) {
	tests := []struct {
		name, data string
		want       map[string]string
	}{
		{
			"basic",
			"be cfbegad cbdgef fgaecd cgeb fdcge agebfd fecdb fabcd edb",
			map[string]string{
				"abcdefg": "8",
				"abcdf":   "2",
				"abdefg":  "0",
				"acdefg":  "6",
				"bcdef":   "3",
				"bcdefg":  "9",
				"bceg":    "4",
				"bde":     "7",
				"be":      "1",
				"cdefg":   "5",
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := demap(tc.data)
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
			testgeneric.GetHandle(t, "testdata/input1.txt"),
			61229,
			false,
		},
		{
			"aocd8",
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
