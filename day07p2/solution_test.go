package day07p2

import (
	"strings"
	"testing"

	"aoc/utils"
)

var testInput = `32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483`

func TestSolve(t *testing.T) {
	tests := []struct {
		input  string
		answer int64
	}{
		{testInput, 5905},
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
