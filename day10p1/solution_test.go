package day10p1

import (
	"strings"
	"testing"

	"aoc/utils"
)

var testInput = `.....
.S-7.
.|.|.
.L-J.
.....`

var testInput2 = `7-F7-
.FJ|7
SJLL7
|F--J
LJ.LJ`

func TestSolve(t *testing.T) {
	tests := []struct {
		input  string
		answer uint64
	}{
		{testInput, 4},
		{testInput2, 8},
	}

	if testing.Verbose() {
		utils.Verbose = true
	}

	for _, test := range tests {
		r := strings.NewReader(test.input)

		result := Solve(r).(uint64)

		if result != test.answer {
			t.Errorf("Expected %d, got %d", test.answer, result)
		}
	}
}
