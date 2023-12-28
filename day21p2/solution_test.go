package day21p2

import (
	"strings"
	"testing"

	"aoc/utils"
)

var testInput = `...........
.....###.#.
.###.##..#.
..#.#...#..
....#.#....
.##..S####.
.##..#...#.
.......##..
.##.#.####.
.##..##.##.
...........`

func TestSolve(t *testing.T) {
	tests := []struct {
		input  string
		steps  int
		answer uint64
	}{
		{testInput, 6, 16},
		{testInput, 10, 50},
		{testInput, 50, 1594},
		{testInput, 100, 6536},
		{testInput, 500, 167004},
		{testInput, 1000, 668697},
		{testInput, 5000, 16733044},
	}

	if testing.Verbose() {
		utils.Verbose = true
	}

	for _, test := range tests {
		r := strings.NewReader(test.input)

		result := solveSteps(r, test.steps)

		if result != test.answer {
			t.Errorf("Expected %d, got %d", test.answer, result)
		}
	}
}
