package main

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/hashicorp/go-multierror"
	"github.com/pkg/errors"
)

type line struct {
	ax, ay, zx, zy int
}

func stringToLine(s string) (line, error) {
	var merr *multierror.Error
	tuple := strings.Fields(s)
	ax, ay, foundA := strings.Cut(tuple[0], ",")
	zx, zy, foundZ := strings.Cut(tuple[2], ",")
	if !foundA || !foundZ {
		return line{}, fmt.Errorf("failed to find coords in (%s)", tuple)
	}
	axn, err := strconv.ParseInt(ax, 10, 64)
	merr = multierror.Append(err, merr)
	ayn, err := strconv.ParseInt(ay, 10, 64)
	merr = multierror.Append(err, merr)
	zxn, err := strconv.ParseInt(zx, 10, 64)
	merr = multierror.Append(err, merr)
	zyn, err := strconv.ParseInt(zy, 10, 64)
	merr = multierror.Append(err, merr)

	err = merr.ErrorOrNil()
	if err != nil {
		return line{}, errors.Wrapf(err, "failed to parse line (%s): %v", s, err)
	}
	return line{int(axn), int(ayn), int(zxn), int(zyn)}, nil
}

func sign(i int) int {
	if i > 0 {
		return 1
	} else if i < 0 {
		return -1
	}
	return 0
}

func expandLine(l line) [][]int {
	diffx := l.zx - l.ax
	diffy := l.zy - l.ay
	diffx, diffy = sign(diffx), sign(diffy)
	lx, ly := l.ax, l.ay
	var out [][]int
	for {
		out = append(out, []int{lx, ly})
		if lx == l.zx && ly == l.zy {
			break
		}
		lx += diffx
		ly += diffy
	}
	return out
}

func p1(r io.Reader) (int, error) {
	s := bufio.NewScanner(r)
	floor := make(map[string]uint)
	for s.Scan() {
		line, err := stringToLine(s.Text())
		if err != nil {
			return 0, err
		}
		points := expandLine(line)
		for _, point := range points {
			floor[fmt.Sprintf("%v-%v", point[0], point[1])]++
		}
	}
	var c int
	for _, v := range floor {
		if v > 1 {
			c++
		}
	}
	return c, nil
}
