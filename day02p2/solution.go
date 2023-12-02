package day02p2

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
		game := parseGame(ln)

		maximums := make(map[string]int)
		for _, draw := range game.Draws {
			for k, v := range draw {
				if v > maximums[k] {
					maximums[k] = v
				}
			}
		}

		sum += maximums["red"] * maximums["green"] * maximums["blue"]
	}

	return sum
}

type Game struct {
	Id    int
	Draws []map[string]int
}

func parseGame(line string) Game {
	firstSplit := strings.Split(line, ": ")

	idstrings := strings.Fields(firstSplit[0])

	result := Game{}

	var err error
	result.Id, err = strconv.Atoi(idstrings[1])
	utils.Check(err, "Unable to convert %s to integer", idstrings[1])

	draws := strings.Split(firstSplit[1], "; ")

	for _, draw := range draws {
		drawmap := make(map[string]int)
		colorsets := strings.Split(draw, ", ")
		for _, colorset := range colorsets {
			parts := strings.Fields(colorset)
			n, err := strconv.Atoi(parts[0])
			utils.Check(err, "Unable to convert %s to integer", parts[0])
			drawmap[parts[1]] = n
		}
		result.Draws = append(result.Draws, drawmap)
	}

	return result
}
