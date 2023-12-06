package day06p2

import (
	"io"
	"math"
	"strconv"
	"strings"

	"aoc/utils"
)

func Solve(r io.Reader) any {
	lines := utils.ReadLines(r)

	timeStrings := strings.Fields(lines[0])
	distanceStrings := strings.Fields(lines[1])

	timeString := strings.Join(timeStrings[1:], "")
	distanceString := strings.Join(distanceStrings[1:], "")

	time, err := strconv.ParseFloat(timeString, 64)
	utils.Check(err, "Unable to parse %s into float", timeString)

	distance, err := strconv.ParseFloat(distanceString, 64)

	tPressMin := 0.5*time - 0.5*math.Sqrt(time*time-4.0*distance)

	tPressMax := 0.5*time + 0.5*math.Sqrt(time*time-4.0*distance)

	winways := math.Floor(tPressMax) - math.Ceil(tPressMin)
	if math.Abs(tPressMin-math.Ceil(tPressMin)) < 1.0e-8 {
		winways -= 1.0
	}

	if math.Abs(tPressMax-math.Floor(tPressMax)) < 1.0e-8 {
		winways -= 1.0
	}

	return int(winways) + 1
}
