package day08p2

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

	cyclesteps := int64(1)

	// Let's find the ones ending in A
	for pos := range neighbors {
		if !strings.HasSuffix(pos, "A") {
			continue
		}
		var steps int64

		hits := []Hit{}

	outer:
		for {
			for i, d := range moves {
				next := neighbors[pos]
				steps++
				switch d {
				case 'L':
					pos = next[0]
				case 'R':
					pos = next[1]
				}

				if strings.HasSuffix(pos, "Z") {
					h := Hit{pos, i, steps}
					for _, ph := range hits {
						if h.Candidate == ph.Candidate {
							// An examination shows that the interval between hits appears to be the same as the
							// time from start to first hit. This is slightly contrived, but makes the solution easier.
							// This is not a general solution, but a solution that works for this and I suspect all
							// advent of code inputs for this day.
							cyclesteps = utils.Lcm(cyclesteps, h.Steps-ph.Steps)
							break outer
						}
					}
					hits = append(hits, h)
				}
			}
		}
	}
	return cyclesteps
}

// State when destination candidate is hit
type Hit struct {
	Candidate string // Ends with Z
	Cycle     int    // i value when hit
	Steps     int64  // step at which it was hit
}
