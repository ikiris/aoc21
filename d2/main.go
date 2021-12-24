package main

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

func p1(r io.Reader) (int64, error) {
	var x, depth int64
	dir := map[string]func(int64){
		"forward": func(i int64) { x += i },
		"up":      func(i int64) { depth -= i },
		"down":    func(i int64) { depth += i },
	}

	s := bufio.NewScanner(r)
	for s.Scan() {
		d, v, ok := strings.Cut(s.Text(), " ")
		if !ok {
			return 0, fmt.Errorf("malformed input string: %s", s.Text())
		}
		i, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return 0, errors.Wrapf(err, "failed to parse %v:", s.Text())
		}
		dir[d](i)
	}
	return depth * x, nil
}

func p2(r io.Reader) (int64, error) {
	var x, depth, aim int64
	dir := map[string]func(int64){
		"forward": func(i int64) {
			x += i
			depth += i * aim
		},
		"up":   func(i int64) { aim -= i },
		"down": func(i int64) { aim += i },
	}

	s := bufio.NewScanner(r)
	for s.Scan() {
		d, v, ok := strings.Cut(s.Text(), " ")
		if !ok {
			return 0, fmt.Errorf("malformed input string: %s", s.Text())
		}
		i, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return 0, errors.Wrapf(err, "failed to parse %v:", s.Text())
		}
		dir[d](i)
	}
	return depth * x, nil
}
