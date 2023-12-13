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
		valid := findValid([]rune(bld.String()), groups)
		if utils.Verbose {
			fmt.Printf("\tValid: %d\n\n", valid)
		}
		sum += valid
	}
	return sum
}

func findValid(row []rune, groups []int) int64 {
	valids := int64(0)

	if isCompleteAndValid(row, groups) {
		valids = 1
		return 1
	}

	for _, r := range findNext(row, groups) {
		valids += findValid(r, groups)
	}

	return valids
}

func findNext(row []rune, groups []int) [][]rune {
	cgroups := []int{}
	inGroup := false
	ng := 0
	firstQ := len(row)
	ret := [][]rune{}
	brokenSum := 0
	unknownSum := 0

	// fmt.Println(string(row))
	for i, r := range row {
		switch r {
		case '#':
			if i < firstQ {
				inGroup = true
				ng++
			}
			brokenSum++
		case '.':
			if i < firstQ {
				inGroup = false
				if ng > 0 {
					cgroups = append(cgroups, ng)
					ng = 0
				}
			}
		case '?':
			if firstQ > i {
				firstQ = i
			}
			unknownSum++
		}
	}
	if firstQ == len(row) {
		return ret
	}

	// Check there are enough remainging to make all groups
	remain := len(row) - firstQ
	if inGroup {
		remain += ng
	}

	needForGroupings := 0
	for i := len(cgroups); i < len(groups); i++ {
		needForGroupings += groups[i] + 1
	}
	needForGroupings--

	if needForGroupings > remain {
		return ret
	}

	sum := 0
	for _, g := range groups {
		sum += g
	}

	if brokenSum+unknownSum < sum {
		// invalid
		return ret
	}

	if len(cgroups) > len(groups) {
		// invalid
		return ret
	}

	if len(cgroups) == len(groups) {
		for i := firstQ; i < len(row); i++ {
			if row[i] == '#' {
				return ret
			}
			row[i] = '.'
		}
		ret = append(ret, row)
		return ret
	}

	// Complete next group of broken springs
	need := groups[len(cgroups)]
	endPoint := len(row)
	if inGroup {
		// group must go here
		need -= ng
		endPoint = firstQ + 1
	} else {
		for i := firstQ + 1; i < len(row); i++ {
			if row[i] == '#' {
				endPoint = i + 1
				break
			}
		}
	}

	if need+firstQ > len(row) {
		// not enough room
		return ret
	}

outer:
	for start := firstQ; start < endPoint; start++ {
		newrow := make([]rune, len(row))
		copy(newrow, row)
		for i := firstQ; i < start; i++ {
			newrow[i] = '.'
		}
		for i := start; i < need+start; i++ {
			if i == len(newrow) || newrow[i] == '.' {
				// wont work
				continue outer
			}
			newrow[i] = '#'
		}
		if need+start < len(row) {
			if row[need+start] == '#' {
				// invalid
				continue outer
			}
			newrow[need+start] = '.'
		}
		ret = append(ret, newrow)
	}

	return ret
}

func isCompleteAndValid(row []rune, groups []int) bool {
	cgroups := []int{}
	ng := 0

	for _, r := range row {
		switch r {
		case '?':
			return false
		case '#':
			ng++
		case '.':
			if ng > 0 {
				cgroups = append(cgroups, ng)
				ng = 0
			}
		}
	}
	if ng > 0 {
		cgroups = append(cgroups, ng)
	}

	if len(groups) != len(cgroups) {
		return false
	}

	for i, ng := range cgroups {
		if ng != groups[i] {
			return false
		}
	}

	// fmt.Printf("\tValid: %s\n", string(row))
	return true
}
