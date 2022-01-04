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

func getNums(r io.Reader) ([]int, int, int, error) {
	min, max := 9999999, 0
	s := bufio.NewScanner(r)
	if !s.Scan() {
		return nil, 0, 0, errors.New("empty buffer!")
	}
	line := strings.Split(s.Text(), ",")
	var nums []int
	for _, str := range line {
		n, err := strconv.Atoi(str)
		if err != nil {
			return nil, 0, 0, errors.Wrapf(err, "failed to parse atoi(%s)", str)
		}
		min = generic.Min(n, min)
		max = generic.Max(n, max)
		nums = append(nums, n)
	}
	return nums, min, max, nil
}

func p1(r io.Reader) (int64, error) {
	nums, _, max, err := getNums(r)
	if err != nil {
		return 0, err
	}

	var dist int64
	// I don't think this is the slow way, but there's probably something faster
	sort.Ints(nums) // put them in order so we can be fancy
	var min int64 = 9999999999
	first := nums[0]
	ln := len(nums)
	for _, n := range nums {
		dist += int64(n - first)
	}

	p := 0
	for i := 1; i < max; i++ { // check all points between min - max
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

func p2(r io.Reader) (int64, error) {
	nums, min, max, err := getNums(r)
	if err != nil {
		return 0, err
	}

	var minDist int64 = 9999999999
	// guess we're doing this the hard way.
	for i := min; i <= max; i++ {
		var dist int64
		for _, n := range nums {
			delta := int64(generic.Abs(n - i))
			dist += (delta * (delta + 1)) / 2
		}
		minDist = generic.Min(dist, minDist)
	}
	return minDist, nil
}
