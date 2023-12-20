package day20p1

import (
	"strings"
	"testing"

	"aoc/utils"
)

var testInput = `broadcaster -> a, b, c
%a -> b
%b -> c
%c -> inv
&inv -> a`

var testInput2 = `broadcaster -> a
%a -> inv, con
&inv -> b
%b -> con
&con -> output`

func TestSolve(t *testing.T) {
	tests := []struct {
		input  string
		answer uint64
	}{
		{testInput, 32000000},
		{testInput2, 11687500},
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
