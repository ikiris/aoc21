package main

import (
	"bufio"
	"io"
	"strconv"

	"github.com/pkg/errors"
)

func p1(r io.Reader) (int64, error) {
	s := bufio.NewScanner(r)
	var c int64
	var last int64
	for s.Scan() {
		i, err := strconv.ParseInt(s.Text(), 10, 64)
		if err != nil {
			return 0, errors.Wrapf(err, "failed to parse %v:", s.Text())
		}
		if i > last {
			c++
		}
		last = i
	}
	return c - 1, nil
}

func p2(r io.Reader) (int64, error) {
	s := bufio.NewScanner(r)
	var c int64
	var a, b int64
	var sum int64
	for s.Scan() {
		i, err := strconv.ParseInt(s.Text(), 10, 64)
		if err != nil {
			return 0, errors.Wrapf(err, "failed to parse %v:", s.Text())
		}
		newsum := a + b + i
		a, b = b, i
		if newsum > sum {
			c++
		}
		sum = newsum
	}
	return c - 3, nil
}
