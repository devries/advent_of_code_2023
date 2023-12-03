package day03p2

import (
	"io"

	"aoc/utils"
)

func Solve(r io.Reader) any {
	lines := utils.ReadLines(r)

	symbols := make(map[utils.Point]rune)

	numbers := []Number{}

	for j, ln := range lines {
		innumber := false

		for i, r := range ln {
			switch {
			case r == '.':
				// do nothing
				innumber = false
			case r >= '0' && r <= '9':
				// number
				switch innumber {
				case false:
					innumber = true
					var newNumber Number
					newNumber.Start = utils.Point{X: i, Y: j}
					newNumber.End = newNumber.Start
					newNumber.Value = int64(r - '0')
					numbers = append(numbers, newNumber)
				case true:
					idx := len(numbers) - 1
					numbers[idx].End = utils.Point{X: i, Y: j}
					numbers[idx].Value = numbers[idx].Value*10 + int64(r-'0')
				}
			default:
				// symbol
				symbols[utils.Point{X: i, Y: j}] = r
				innumber = false
			}
		}
	}

	var sum int64

	gearAdjacents := make(map[utils.Point][]int64)

	for _, n := range numbers {
		surroundingPoints := pointsAround(n.Start, n.End)
		for _, p := range surroundingPoints {
			if symbols[p] == '*' {
				gearAdjacents[p] = append(gearAdjacents[p], n.Value)
				break
			}
		}
	}

	for _, nums := range gearAdjacents {
		if len(nums) == 2 {
			sum += nums[0] * nums[1]
		}
	}

	return sum
}

type Number struct {
	Value int64
	Start utils.Point
	End   utils.Point
}

func pointsAround(start utils.Point, end utils.Point) []utils.Point {
	yupper := start.Y - 1
	ylower := start.Y + 1
	yline := start.Y

	xmin := start.X - 1
	xmax := end.X + 1

	result := []utils.Point{{X: xmin, Y: yline}, {X: xmax, Y: yline}}

	for i := xmin; i <= xmax; i++ {
		result = append(result, utils.Point{X: i, Y: yupper})
		result = append(result, utils.Point{X: i, Y: ylower})
	}

	return result
}
