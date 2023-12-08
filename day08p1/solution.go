package day08p1

import (
	"io"
	"strings"

	"aoc/utils"
)

func Solve(r io.Reader) any {
	lines := utils.ReadLines(r)

	// Get all LR moves
	moves := []rune(lines[0])

	neighbors := make(map[string][]string)
	for i := 2; i < len(lines); i++ {
		parts := strings.Split(lines[i], " = ")

		destinations := strings.Trim(parts[1], "()")

		leftright := strings.Split(destinations, ", ")

		neighbors[parts[0]] = leftright
	}

	next := neighbors["AAA"]
	steps := 0

outer:
	for {
		for _, d := range moves {
			steps++
			switch d {
			case 'L':
				if next[0] == "ZZZ" {
					break outer
				}
				next = neighbors[next[0]]
			case 'R':
				if next[1] == "ZZZ" {
					break outer
				}
				next = neighbors[next[1]]
			}
		}
	}

	return steps
}
