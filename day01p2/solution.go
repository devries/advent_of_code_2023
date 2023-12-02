package day01p2

import (
	"io"
	"strings"

	"aoc/utils"
)

func Solve(r io.Reader) any {
	lines := utils.ReadLines(r)

	matchmap := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
		"0":     0,
		"1":     1,
		"2":     2,
		"3":     3,
		"4":     4,
		"5":     5,
		"6":     6,
		"7":     7,
		"8":     8,
		"9":     9,
	}

	sum := 0

	for _, ln := range lines {
		first := -1
		last := -1
		minpos := len(ln) + 1
		maxpos := -1

		for k, v := range matchmap {
			p := strings.Index(ln, k)
			if p == -1 {
				// substring not found
				continue
			}
			if p < minpos {
				first = v
				minpos = p
			}

			p = strings.LastIndex(ln, k)
			if p > maxpos {
				last = v
				maxpos = p
			}
		}
		sum += 10*first + last
	}
	return sum
}
