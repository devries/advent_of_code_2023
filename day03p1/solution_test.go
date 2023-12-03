package day03p1

import (
	"strings"
	"testing"

	"aoc/utils"
)

var testInput = `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`

func TestSolve(t *testing.T) {
	tests := []struct {
		input  string
		answer int64
	}{
		{testInput, 4361},
	}

	if testing.Verbose() {
		utils.Verbose = true
	}

	for _, test := range tests {
		r := strings.NewReader(test.input)

		result := Solve(r).(int64)

		if result != test.answer {
			t.Errorf("Expected %d, got %d", test.answer, result)
		}
	}
}
