package day17p1

import (
	"strings"
	"testing"

	"aoc/utils"
)

var testInput = `2413432311323
3215453535623
3255245654254
3446585845452
4546657867536
1438598798454
4457876987766
3637877979653
4654967986887
4564679986453
1224686865563
2546548887735
4322674655533`

func TestSolve(t *testing.T) {
	tests := []struct {
		input  string
		answer uint64
	}{
		{testInput, 102},
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
