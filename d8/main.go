package main

import (
	"bufio"
	"errors"
	"io"
	"strconv"
	"strings"

	"github.com/ikiris/aoc21/generic"
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

func demap(scramble string) map[string]string {
	scrambleN := strings.Fields(scramble)
	histogram := make(map[rune]int)
	vmap := map[int]string{42: "0", 17: "1", 34: "2", 39: "3", 30: "4", 37: "5", 41: "6", 25: "7", 49: "8", 45: "9"}

	// get the substitution cypher character histogram
	for _, s := range scrambleN {
		for _, r := range []rune(s) {
			histogram[r]++
		}
	}

	dt := make(map[string]string)
	for _, s := range scrambleN {
		ttl := 0
		for _, r := range []rune(s) {
			ttl += histogram[r]
		}
		ss := generic.SortString(s)
		dt[ss] = vmap[ttl]
	}
	return dt
}

func p2(r io.Reader) (int64, error) {
	s := bufio.NewScanner(r)
	var count int64
	for s.Scan() {
		scramble, output, ok := strings.Cut(s.Text(), "|")
		if !ok {
			return 0, errors.New("shrug")
		}

		//get the mapping table for the scramble string
		dmap := demap(scramble)

		// ok we know what's what, so lets convert + add them.
		var num string
		digits := strings.Fields(output)
		for _, d := range digits {
			num += dmap[generic.SortString(d)]
		}
		cNum, err := strconv.Atoi(num)
		if err != nil {
			return 0, err
		}
		count += int64(cNum)
	}
	return count, nil
}
