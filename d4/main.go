package main

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

func parseCalls(s string) ([]uint, error) {
	c := strings.Split(s, ",")
	i := make([]uint, len(c))
	for k, v := range c {
		iv, err := strconv.ParseUint(v, 10, 64)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to parse(%v)", v)
		}
		i[k] = uint(iv)
	}
	return i, nil
}

type board struct {
	b   []uint
	m   map[uint]uint
	won bool
}

func parseBoard(s []string) (*board, error) {
	b := &board{
		b: make([]uint, 0, 25),
		m: make(map[uint]uint),
	}
	for _, s := range s {
		for _, v := range strings.Fields(s) {
			vi, err := strconv.ParseUint(v, 10, 64)
			if err != nil {
				return nil, errors.Wrapf(err, "failed to parse %v", v)
			}
			b.b = append(b.b, uint(vi))
			b.m[uint(vi)] = uint(len(b.b)) - 1
		}
	}
	return b, nil
}

func parseBoards(s *bufio.Scanner) ([]*board, error) {
	b := make([]string, 0, 5)
	var boards []*board
	for s.Scan() {
		if s.Text() == "" {
			continue
		}
		b = append(b, s.Text())
		if len(b) == 5 {
			board, err := parseBoard(b)
			if err != nil {
				return nil, errors.Wrapf(err, "failed parsing board(%v)", board)
			}
			boards = append(boards, board)
			b = make([]string, 0, 5)
		}
	}
	return boards, nil
}

func p1(r io.Reader) (uint64, error) {
	s := bufio.NewScanner(r)
	if !s.Scan() {
		return 0, fmt.Errorf("empty buf!")
	}

	calls, err := parseCalls(s.Text())
	if err != nil {
		return 0, errors.Wrap(err, "failed to parse call line")
	}

	boards, err := parseBoards(s)

	for _, c := range calls {
		for _, board := range boards {
			if won := board.Call(c); won {
				return winnerP1(board, c)
			}
		}
	}
	return 0, nil
}

func p2(r io.Reader) (uint64, error) {
	s := bufio.NewScanner(r)
	if !s.Scan() {
		return 0, fmt.Errorf("empty buf!")
	}

	calls, err := parseCalls(s.Text())
	if err != nil {
		return 0, errors.Wrap(err, "failed to parse call line")
	}

	boards, err := parseBoards(s)

	var winners []*board
	var lastCall uint
	for _, c := range calls {
		for _, board := range boards {
			if won := board.Call(c); won {
				lastCall = c
				winners = append(winners, board)
			}
		}
	}
	return winnerP1(winners[len(winners)-1], lastCall)
}

func (b *board) Call(c uint) bool {
	// do we have it?
	l, ok := b.m[c]
	if !ok || b.won { // or have we already won? - Lazyyyyyy
		return false
	}

	// Blow the fuse for the position.
	b.b[l] = 255

	// Did this win?
	col := l % 5
	row := l - col
	rs := b.b[row] + b.b[row+1] + b.b[row+2] + b.b[row+3] + b.b[row+4]
	cs := b.b[col] + b.b[col+5] + b.b[col+10] + b.b[col+15] + b.b[col+20]
	if rs == 1275 || cs == 1275 {
		b.won = true
		return true
	}
	return false
}

func winnerP1(b *board, c uint) (uint64, error) {
	var s uint64
	for _, v := range b.b {
		if v >= 255 {
			continue
		}
		s += uint64(v)
	}
	return s * uint64(c), nil
}
