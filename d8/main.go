package main

import (
	"bufio"
	"errors"
	"io"
	"strings"
)

func p1(r io.Reader) (int, error) {
	s := bufio.NewScanner(r)
	count := 0
	for s.Scan() {
		_, output, ok := strings.Cut(s.Text(), "|")
		if !ok {
			return 0, errors.New("shrug")
		}
		digits := strings.Fields(output)
		for _, d := range digits {
			switch len(d) {
			case 2, 4, 3, 7:
				count++
			}
		}
	}
	return count, nil
}
