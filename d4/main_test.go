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

var (
	throwboard1 = board{
		b: []uint{24, 66, 11, 44, 51, 32, 37, 83, 69, 85, 46, 59, 14, 99, 76, 71, 28, 94, 35, 98, 16, 40, 74, 80, 6},
		m: map[uint]uint{
			24: 0,
			66: 1,
			11: 2,
			44: 3,
			51: 4,
			32: 5,
			37: 6,
			83: 7,
			69: 8,
			85: 9,
			46: 10,
			59: 11,
			14: 12,
			99: 13,
			76: 14,
			71: 15,
			28: 16,
			94: 17,
			35: 18,
			98: 19,
			16: 20,
			40: 21,
			74: 22,
			80: 23,
			6:  24,
		},
	}
	throwparse1 = []string{
		"24 66 11 44 51",
		"32 37 83 69 85",
		"46 59 14 99 76",
		"71 28 94 35 98",
		"16 40 74 80  6",
	}
	throwparse2 = []string{
		"3 15  0  2 22",
		"9 18 13 17  5",
		"19  8  7 25 23",
		"20 11 10 24  4",
		"14 21 16 12  6",
	}
)

func TestParseBoard(t *testing.T) {
	tests := []struct {
		name    string
		data    []string
		want    *board
		wantErr bool
	}{
		{
			"Basic",
			throwparse1,
			&throwboard1,
			false,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := parseBoard(tc.data)
			if (err != nil) != tc.wantErr {
				t.Errorf("func %s goterr: %v wanted: %v", tc.name, err, tc.wantErr)
			}
			if diff := cmp.Diff(tc.want, got, cmp.AllowUnexported(board{})); diff != "" {
				t.Errorf("func diff(-got,+want):%v", diff)
			}
		})
	}
}

func TestCall(t *testing.T) {
	tests := []struct {
		name      string
		boards    [][]string
		calls     []uint
		want      bool
		wantBoard int
		winCall   uint
	}{
		{
			"Horizontal Basic Win",
			[][]string{
				throwparse1,
				throwparse2,
			},
			[]uint{9, 18, 13, 17, 5},
			true,
			1,
			5,
		},
		{
			"Vertical Basic Win",
			[][]string{
				throwparse1,
				throwparse2,
			},
			[]uint{15, 18, 8, 11, 21},
			true,
			1,
			21,
		},
		{
			"NoJoy",
			[][]string{
				throwparse1,
				throwparse2,
			},
			[]uint{0, 11, 17, 13, 19},
			false,
			0,
			0,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			var boards []*board
			for _, b := range tc.boards {
				bp, err := parseBoard(b)
				if err != nil {
					t.Fatalf("failed to parse board %v: %v", b, err)
				}
				boards = append(boards, bp)
			}

			var gotB int
			var gotC uint
			var got bool
		outer:
			for _, c := range tc.calls {
				for bi, board := range boards {
					if won := board.Call(c); won {
						gotB, gotC, got = bi, c, won
						break outer
					}
				}
			}

			if gotB != tc.wantBoard {
				t.Errorf("func got winner %d wanted %d", gotB, tc.wantBoard)
			}

			if gotC != tc.winCall {
				t.Errorf("func got winner call %d wanted %d", gotC, tc.winCall)
			}

			if got != tc.want {
				t.Errorf("func got: %v, want: %v", got, tc.want)
			}
		})
	}
}

func TestP1(t *testing.T) {
	tests := []struct {
		name    string
		data    io.Reader
		want    uint64
		wantErr bool
	}{
		{
			"basic",
			getHandle(t, "testdata/input1.txt"),
			4512,
			false,
		}, {
			"aocd4",
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

func TestP2(t *testing.T) {
	tests := []struct {
		name    string
		data    io.Reader
		want    uint64
		wantErr bool
	}{
		{
			"basic",
			getHandle(t, "testdata/input1.txt"),
			1924,
			false,
		}, {
			"aocd4",
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
