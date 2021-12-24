package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
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
	log.Println(s.Text())
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
