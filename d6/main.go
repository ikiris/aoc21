package main

import (
	"bufio"
	"errors"
	"io"
	"strconv"
	"strings"

	"github.com/ikiris/aoc21/generic"
)

func p1(r io.Reader, days int) (int64, error) {
	s := bufio.NewScanner(r)
	if !s.Scan() {
		return 0, errors.New("empty buffer!")
	}
	line := strings.Split(s.Text(), ",")
	fish := make(map[int]int64)
	for _, val := range line {
		v, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			return 0, err
		}
		fish[int(v)]++
	}
	for i := 0; i < days; i++ {
		newfish := make(map[int]int64)
		for age := 0; age <= 8; age++ {
			newage := age - 1
			if age == 0 {
				newage = 6
				newfish[8] += fish[age]
			}
			newfish[newage] += fish[age]
		}
		fish = newfish
	}
	return generic.AddMapVals(fish), nil
}
