package day04p2

import (
	"io"
	"strconv"
	"strings"

	"aoc/utils"
)

func Solve(r io.Reader) any {
	lines := utils.ReadLines(r)

	// this will start at 0 and we will add to it, so the multiplier is this plus 1
	multiplier := make(map[int]int)
	max := 0

	for i, ln := range lines {
		card := parseLine(ln)

		winners := make(map[int]bool)
		for _, v := range card.Winners {
			winners[v] = true
		}

		matches := 0

		for _, v := range card.Selected {
			if winners[v] {
				matches++
			}
		}

		for j := i + 1; j <= i+matches; j++ {
			// Add cards equal to the number of copies of this card you have
			multiplier[j] = multiplier[j] + multiplier[i] + 1
		}
		max = i
	}

	sum := 0

	for i := 0; i <= max; i++ {
		sum += multiplier[i] + 1
	}
	return sum
}

type Card struct {
	Winners  []int
	Selected []int
}

func parseLine(ln string) Card {
	titleSplit := strings.Split(ln, ": ")

	contents := strings.Split(titleSplit[1], " | ")

	return Card{stringToNumbers(contents[0]), stringToNumbers(contents[1])}
}

func stringToNumbers(s string) []int {
	groups := strings.Fields(s)

	var result []int

	for _, v := range groups {
		n, err := strconv.Atoi(v)
		utils.Check(err, "unable to convert %s to integer", v)

		result = append(result, n)
	}

	return result
}
