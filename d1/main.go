package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"

	"github.com/pkg/errors"
)

func main() {
	fh, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("failed to open: %v", err)
	}

	c, err := doThings(fh)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(c)
}

func doThings(r io.Reader) (int64, error) {
	s := bufio.NewScanner(r)
	var last int64
	var c int64
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
	return c, nil
}
