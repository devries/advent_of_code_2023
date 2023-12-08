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
			for _, d := range moves {
				next := neighbors[pos]
				steps++
				switch d {
				case 'L':
					pos = next[0]
				case 'R':
					pos = next[1]
				}

				if strings.HasSuffix(pos, "Z") {
					h := Hit{pos, steps}
					for _, ph := range hits {
						if h.Candidate == ph.Candidate && h.Steps == 2*ph.Steps {
							// Let's check if we are getting back to the same position in regular cycles.
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
	Steps     int64  // step at which it was hit
}
