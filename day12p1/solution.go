package day12p1

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

		groups := make([]int64, len(numbers))

		var err error
		for i, v := range numbers {
			groups[i], err = strconv.ParseInt(v, 10, 64)
			utils.Check(err, "Unable to convert %s to int64", v)
		}

		sum += findValid([]rune(parts[0]), groups)
	}
	return sum
}

func findValid(row []rune, groups []int64) int64 {
	valids := int64(0)

	for i, r := range row {
		if r == '?' {
			// replace with functional spring
			trialA := make([]rune, len(row))
			copy(trialA, row)
			trialA[i] = '.'
			if isValid(trialA, groups) {
				valids += findValid(trialA, groups)
			}

			// replace with broken sprint
			trialB := make([]rune, len(row))
			copy(trialB, row)
			trialB[i] = '#'
			if isValid(trialB, groups) {
				valids += findValid(trialB, groups)
			}

			return valids
		}
	}

	if isValid(row, groups) {
		valids = 1
		if utils.Verbose {
			fmt.Printf("\t%s\n", string(row))
		}
	}
	return valids
}

func isValid(row []rune, groups []int64) bool {
	cgroups := []int64{}

	n := int64(0)

	for _, r := range row {
		switch r {
		case '#':
			n++
		case '.':
			if n > 0 {
				cgroups = append(cgroups, n)
				n = 0
			}
		case '?':
			return true
		}
	}
	if n > 0 {
		cgroups = append(cgroups, n)
	}
	n = 0

	if len(cgroups) != len(groups) {
		return false
	}

	for i, g := range groups {
		if cgroups[i] != g {
			return false
		}
	}

	return true
}
