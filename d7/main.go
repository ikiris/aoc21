package main

import (
	"bufio"
	"io"
	"sort"
	"strconv"
	"strings"

	"github.com/ikiris/aoc21/generic"
	"github.com/pkg/errors"
)

func p1(r io.Reader) (int64, error) {
	s := bufio.NewScanner(r)
	if !s.Scan() {
		return 0, errors.New("empty buffer!")
	}
	line := strings.Split(s.Text(), ",")
	var nums []int
	for _, str := range line {
		n, err := strconv.Atoi(str)
		if err != nil {
			return 0, errors.Wrapf(err, "failed to parse atoi(%s)", str)
		}
		nums = append(nums, n)
	}

	var dist int64
	// I don't think this is the slow way, but there's probably something faster
	sort.Ints(nums) // put them in order so we can be fancy
	var min int64
	first := nums[0]
	ln := len(nums)
	for _, n := range nums {
		dist += int64(n - first)
	}
	min = dist
	p := 0
	for i := 1; i < nums[ln-1]; i++ { // check all points between min - max
		v := i + first
		for nums[p] < v { // move the array pointer as we pass them.
			p++
		}
		dist += int64(p)
		dist -= int64(ln - p)
		min = generic.Min(dist, min)
	}
	return min, nil
}
