package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"strconv"

	"github.com/pkg/errors"
)

func p1(r io.Reader) (int64, error) {
	s := bufio.NewScanner(r)
	var c float64
	if !s.Scan() {
		return 0, fmt.Errorf("empty buf!")
	}
	n := make([]int64, len(s.Text()))
	cap := len(n)
	base := cap - 1
	done := true
	for done {
		i, err := strconv.ParseInt(s.Text(), 2, 32)
		if err != nil {
			return 0, errors.Wrapf(err, "failed to parse %v:", s.Text())
		}
		for ni := base; ni >= 0; ni-- {
			iv := 1 << (base - ni)
			n[ni] += i & int64(iv)
		}
		c++
		done = s.Scan()
	}
	var g int64
	for n, i := range n {
		f := i >> (base - n)
		ro := math.Round(float64(f) / c)
		g += int64(ro) << (base - n)
	}
	e := ^g & ((1 << base) - 1)
	return g * e, nil
}

func zipfilter(intList []int64, pos int, comparable func(a, b int) bool) ([]int64, error) {
	if len(intList) <= 1 {
		return intList, nil
	}
	if pos < 0 {
		return intList, fmt.Errorf("Negative pos! (%d)", pos)
	}
	var a, b []int64
	for _, v := range intList {
		bit := (v >> pos) & 1
		if bit == 1 {
			a = append(a, v)
		} else {
			b = append(b, v)
		}
	}
	if comparable(len(a), len(b)) {
		return zipfilter(a, pos-1, comparable)
	} else {
		return zipfilter(b, pos-1, comparable)
	}
}

func p2(r io.Reader) (int64, error) {
	comp := map[string]func(a, b int) bool{
		"one":  func(a, b int) bool { return a >= b },
		"zero": func(a, b int) bool { return a < b },
	}
	s := bufio.NewScanner(r)
	if !s.Scan() {
		return 0, fmt.Errorf("empty buf!")
	}
	bitSize := len(s.Text()) - 1
	var a []int64
	v, err := strconv.ParseInt(s.Text(), 2, 64)
	if err != nil {
		return 0, errors.Wrapf(err, "failed to parse %v:", s.Text())
	}
	a = append(a, v)
	// I'm too tired to do this anything other than the memory hog way.
	for s.Scan() {
		v, err := strconv.ParseInt(s.Text(), 2, 64)
		if err != nil {
			return 0, errors.Wrapf(err, "failed to parse %v:", s.Text())
		}
		a = append(a, v)
	}
	oxy, err := zipfilter(a, bitSize, comp["one"])
	if err != nil {
		return 0, err
	}
	co2, err := zipfilter(a, bitSize, comp["zero"])
	if err != nil {
		return 0, err
	}
	return oxy[0] * co2[0], nil
}
