package day15p1

import (
	"io"
	"strings"

	"aoc/utils"
)

func Solve(r io.Reader) any {
	lines := utils.ReadLines(r)

	steps := strings.Split(lines[0], ",")

	var sum uint64
	for _, s := range steps {
		sum += uint64(elfHash(s))
	}
	return sum
}

func elfHash(s string) byte {
	var res byte

	for _, r := range s {
		res += byte(r)
		res *= 17
	}

	return res
}
