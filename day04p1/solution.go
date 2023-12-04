package day04p1

import (
	"io"
	"strconv"
	"strings"

	"aoc/utils"
)

func Solve(r io.Reader) any {
	lines := utils.ReadLines(r)

	sum := 0

	for _, ln := range lines {
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

		switch matches {
		case 0:
			// do nothing
		default:
			sum += 1 << (matches - 1)
		}
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
