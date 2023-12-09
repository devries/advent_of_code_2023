package day09p1

import (
	"strings"
	"testing"

	"aoc/utils"
)

var testInput = `0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45`

func TestSolve(t *testing.T) {
	tests := []struct {
		input  string
		answer int
	}{
		{testInput, 114},
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
