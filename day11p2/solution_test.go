package day11p2

import (
	"strings"
	"testing"

	"aoc/utils"
)

var testInput = `...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#.....`

func TestSolve(t *testing.T) {
	tests := []struct {
		input  string
		factor int64
		answer int64
	}{
		{testInput, 10, 1030},
		{testInput, 100, 8410},
	}

	if testing.Verbose() {
		utils.Verbose = true
	}

	for _, test := range tests {
		r := strings.NewReader(test.input)

		result := SolveFactor(r, test.factor)

		if result != test.answer {
			t.Errorf("Expected %d, got %d", test.answer, result)
		}
	}
}
