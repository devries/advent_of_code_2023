package day15p2

import (
	"strings"
	"testing"

	"aoc/utils"
)

var testInput = `rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7`

func TestSolve(t *testing.T) {
	tests := []struct {
		input  string
		answer uint64
	}{
		{testInput, 145},
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

func TestHash(t *testing.T) {
	tests := []struct {
		input  string
		answer byte
	}{
		{"HASH", 52},
	}

	if testing.Verbose() {
		utils.Verbose = true
	}

	for _, test := range tests {
		result := elfHash(test.input)

		if result != test.answer {
			t.Errorf("Expected %d, got %d", test.answer, result)
		}
	}
}
