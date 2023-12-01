package day01p1

import (
	"strings"
	"testing"

	"aoc/utils"
)

var testInput = `1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet
`

func TestSolve(t *testing.T) {
	tests := []struct {
		input  string
		answer int
	}{
		{testInput, 142},
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
