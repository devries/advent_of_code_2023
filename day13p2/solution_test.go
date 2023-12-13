package day13p2

import (
	"strings"
	"testing"

	"aoc/utils"
)

var testInput = `#.##..##.
..#.##.#.
##......#
##......#
..#.##.#.
..##..##.
#.#.##.#.

#...##..#
#....#..#
..##..###
#####.##.
#####.##.
..##..###
#....#..#`

var testInput2 = `###.#####.#
###.#####.#
#..##.####.
####.#...#.
#.....#..##
.##.#..##.#
##.##.#...#
##.##.#...#
.#..#..##.#
#.....#..##
####.#...#.`

var testInput3 = `..##...
..#.#..
..#.#..
..##...
.#..###
##....#
....#..
#..#...
.##....
...#...
.####..`

func TestSolve(t *testing.T) {
	tests := []struct {
		input  string
		answer int
	}{
		{testInput, 400},
		{testInput2, 700},
		{testInput3, 6},
	}

	if testing.Verbose() {
		utils.Verbose = true
	}

	for _, test := range tests {
		r := strings.NewReader(test.input)

		result := Solve(r).(int)

		if result != test.answer {
			t.Errorf("Expected %d, got %d", test.answer, result)
		}
	}
}
