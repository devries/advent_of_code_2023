package day06p1

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

	sum := 1

	for i := 1; i < len(timeStrings); i++ {
		time, err := strconv.ParseFloat(timeStrings[i], 64)
		utils.Check(err, "Unable to parse %s into float64", timeStrings[i])

		distance, err := strconv.ParseFloat(distanceStrings[i], 64)
		utils.Check(err, "unable to convert %s to float64", distanceStrings[i])

		tPressMin := 0.5*time - 0.5*math.Sqrt(time*time-4.0*distance)

		tPressMax := 0.5*time + 0.5*math.Sqrt(time*time-4.0*distance)

		winways := math.Floor(tPressMax) - math.Ceil(tPressMin)
		if math.Abs(tPressMin-math.Ceil(tPressMin)) < 1.0e-8 {
			winways -= 1.0
		}

		if math.Abs(tPressMax-math.Floor(tPressMax)) < 1.0e-8 {
			winways -= 1.0
		}

		sum *= int(winways) + 1
	}
	return sum
}
