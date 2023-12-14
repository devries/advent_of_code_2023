package day14p1

import (
	"io"

	"aoc/utils"
)

func Solve(r io.Reader) any {
	lines := utils.ReadLines(r)
	ymax := len(lines)
	xmax := len(lines[0])

	pos := make(map[utils.Point]rune)

	for j, ln := range lines {
		for i, r := range ln {
			if r != '.' {
				p := utils.Point{X: i, Y: ymax - j - 1}
				pos[p] = r
			}
		}
	}

	platform := Platform{xmax, ymax, pos}

	tilt(&platform)

	// measure load
	sum := 0
	for k, v := range platform.Positions {
		if v == 'O' {
			sum += k.Y + 1
		}
	}
	return sum
}

type Platform struct {
	XMax      int
	YMax      int
	Positions map[utils.Point]rune
}

func tilt(platform *Platform) {
	// iterate over each row

	for i := 0; i < platform.XMax; i++ {
		blocker := platform.YMax // what blocks the stones
		for j := platform.YMax - 1; j >= 0; j-- {
			pt := utils.Point{X: i, Y: j}
			r := platform.Positions[pt]
			switch r {
			case 'O':
				newPt := utils.Point{X: i, Y: blocker - 1}
				platform.Positions[pt] = 0
				platform.Positions[newPt] = 'O'
				blocker = blocker - 1
			case '#':
				blocker = j
			}
		}
	}
}
