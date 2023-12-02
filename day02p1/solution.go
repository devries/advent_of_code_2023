package day02p1

import (
	"io"
	"strconv"
	"strings"

	"aoc/utils"
)

var maxred = 12
var maxgreen = 13
var maxblue = 14

func Solve(r io.Reader) any {
	lines := utils.ReadLines(r)

	sum := 0

outer:
	for _, ln := range lines {
		game := parseGame(ln)

		for _, draw := range game.Draws {
			if draw["red"] > maxred || draw["green"] > maxgreen || draw["blue"] > maxblue {
				continue outer
			}
		}
		sum += game.Id
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
