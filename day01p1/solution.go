package day01p1

import (
	"io"

	"aoc/utils"
)

func Solve(r io.Reader) any {
	lines := utils.ReadLines(r)

	var first, last int
	sum := 0

	for _, ln := range lines {
		first = -1
		last = -1
		for _, r := range ln {
			if r >= '0' && r <= '9' {
				if first == -1 {
					first = int(r - '0')
				}
				last = int(r - '0')
			}
		}
		sum += 10*first + last
	}
	return sum
}
