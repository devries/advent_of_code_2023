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

		groups := make([]int64, len(numbers)*5)

		for i, v := range numbers {
			n, err := strconv.ParseInt(v, 10, 64)
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
		valid := findValid([]rune(bld.String()), groups)
		if utils.Verbose {
			fmt.Printf("\tValid: %d\n\n", valid)
		}
		sum += valid
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
		// if utils.Verbose {
		// 	fmt.Printf("\t%s\n", string(row))
		// }
	}
	return valids
}

func isValid(row []rune, groups []int64) bool {
	cgroups := []int64{}

	n := int64(0)
	sum := int64(0)
	empties := int64(0)
	firstEmpty := false

	// To add:
	// find minimum possible groups
	// by checking for '.' between remaining '#'
	// also: if !firstEmpty sum==sum(groups)

	for _, r := range row {
		switch r {
		case '#':
			n++
			sum++
		case '.':
			if !firstEmpty && n > 0 {
				cgroups = append(cgroups, n)
				n = 0
			}
		case '?':
			empties++
			if !firstEmpty && n > 0 {
				cgroups = append(cgroups, n)
			}
			n = 0
			firstEmpty = true
		}
	}
	if !firstEmpty && n > 0 {
		cgroups = append(cgroups, n)
	}
	n = 0

	if !firstEmpty && len(cgroups) != len(groups) {
		return false
	}

	if len(cgroups) > len(groups) {
		// more groups already made than allowed
		return false
	}

	totalBroken := int64(0)
	cgl := len(cgroups)
	for i, g := range groups {
		if i < cgl-1 && cgroups[i] != g {
			return false
		} else if i == cgl-1 && cgroups[i] > g {
			// What if there are no empties? then cgroups[i] should equal g
			return false
		}
		totalBroken += g
	}

	if sum > totalBroken {
		// too many defects
		return false
	}

	if sum+empties < totalBroken {
		// not enough to make total broken
		return false
	}

	return true
}
