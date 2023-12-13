package day13p1

import (
	"io"

	"aoc/utils"
)

func Solve(r io.Reader) any {
	lines := utils.ReadLines(r)
	start := 0

	sum := 0
	for i, v := range lines {
		if v == "" || i == len(lines)-1 {
			// end of map parse it
			field := parseField(lines[start:i])

			// Look for vertical symmetry
			v := findSymmetryAxis(field.Verticals)
			sum += v
			// Look for horizontal symmetry
			h := findSymmetryAxis(field.Horizontals)
			sum += 100 * h

			// Move to the next one
			start = i + 1
		}
	}

	return sum
}

type Field struct {
	Horizontals []uint64
	Verticals   []uint64
}

func parseField(lines []string) Field {
	hz := []uint64{}
	vz := []uint64{}

	for j, ln := range lines {
		var hval uint64
		for i, r := range ln {
			var c uint64
			if r == '#' {
				c = 1
			}

			// set horizontal bit
			hval |= c << i
			if j == 0 {
				vz = append(vz, c)
			} else {
				vz[i] |= c << j
			}
		}
		hz = append(hz, hval)
	}

	return Field{hz, vz}
}

func findSymmetryAxis(lines []uint64) int {
	for i := 0; i < len(lines)-1; i++ { // i+1 = lines to left
		for c := 0; c < len(lines)/2+1; c++ {
			l := i - c
			r := i + c + 1

			if l < 0 || r >= len(lines) {
				return i + 1
			}

			if lines[l] != lines[r] {
				break
			}
		}
	}

	return 0
}
