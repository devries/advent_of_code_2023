package day12p2

import (
	"fmt"
	"io"
	"strconv"
	"strings"

	"aoc/utils"
)

func Solve(r io.Reader) any {
	lines := utils.ReadLines(r)

	sum := int64(0)
	for _, ln := range lines {
		if utils.Verbose {
			fmt.Printf("Starting %s\n", ln)
		}

		parts := strings.Fields(ln)

		numbers := strings.Split(parts[1], ",")

		groups := make([]int, len(numbers)*5)

		for i, v := range numbers {
			n, err := strconv.Atoi(v)
			utils.Check(err, "Unable to convert %s to int64", v)
			for j := 0; j < 5; j++ {
				groups[i+len(numbers)*j] = n
			}
		}

		var bld strings.Builder
		for i := 0; i < 5; i++ {
			bld.WriteString(parts[0])
			if i < 4 {
				bld.WriteString("?")
			}
		}

		if utils.Verbose {
			fmt.Printf("\texpanded: %s\n", bld.String())
		}

		s := NewSequence([]rune(bld.String()), groups)
		state := State{0, 0}
		valid := s.CountValid(state)
		if utils.Verbose {
			fmt.Printf("\tValid: %d\n\n", valid)
		}
		sum += valid
	}
	return sum
}

// Memoize based on state
type State struct {
	Start      int
	GroupsDone int
}

// Structure to hold problem and cache
type Sequence struct {
	Row       []rune
	Groups    []int
	NGroups   int             // number of groups
	RowLength int             // length of row
	Seen      map[State]int64 // Previously seen results
}

func NewSequence(row []rune, groups []int) *Sequence {
	s := Sequence{row, groups, len(groups), len(row), make(map[State]int64)}

	return &s
}

func (s *Sequence) CountValid(state State) int64 {
	if v, ok := s.Seen[state]; ok {
		return v
	}

	// Check if there are no more groups to do
	if state.GroupsDone >= s.NGroups {
		valid := true
		for i := state.Start; i < s.RowLength; i++ {
			if s.Row[i] == '#' {
				valid = false
				break
			}
		}
		if valid {
			s.Seen[state] = 1
			return 1
		} else {
			s.Seen[state] = 0
			return 0
		}
	}

	validSolutions := int64(0)
	remainingSprings := s.RowLength - state.Start

	// The number of springs that need to have a specific value are
	// the remaining group counts plus a spacer between each group
	accountedSprings := 0
	for i := state.GroupsDone; i < s.NGroups; i++ {
		accountedSprings += s.Groups[i]
		accountedSprings++ // add spacer
	}
	accountedSprings-- // remove spacer for last one

	// These are springs that are not spacers or broken springs in the rest of the
	// problem
	extraWorkingSprings := remainingSprings - accountedSprings

	//
	for buffer := 0; buffer < extraWorkingSprings+1; buffer++ {
		valid := true
		lastPosition := state.Start + buffer + s.Groups[state.GroupsDone]
		for i := state.Start; i < lastPosition; i++ {
			if i < state.Start+buffer && s.Row[i] == '#' {
				valid = false
				break
			}
			if i >= state.Start+buffer && s.Row[i] == '.' {
				valid = false
				break
			}
		}
		if lastPosition < s.RowLength && s.Row[lastPosition] == '#' {
			valid = false
		}

		if valid {
			newState := State{lastPosition + 1, state.GroupsDone + 1}
			validSolutions += s.CountValid(newState)
		}
	}

	s.Seen[state] = validSolutions
	return validSolutions
}
