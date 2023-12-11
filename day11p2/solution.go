package day11p2

import (
	"io"

	"aoc/utils"

	"github.com/devries/combs"
)

var expansionFactor int64 = 1000000

func Solve(r io.Reader) any {
	return SolveFactor(r, expansionFactor)
}

func SolveFactor(r io.Reader, factor int64) int64 {
	lines := utils.ReadLines(r)

	galaxyRows := make(map[int]bool)
	galaxyColumns := make(map[int]bool)

	galaxies := []utils.Point{}

	// Parse input
	for j, ln := range lines {
		for i, r := range ln {
			if r == '#' {
				p := utils.Point{X: i, Y: j}
				galaxies = append(galaxies, p)
				galaxyRows[j] = true
				galaxyColumns[i] = true
			}
		}
	}

	sum := int64(0)

	// go through all combinations

	for combo := range combs.Combinations(2, galaxies) {
		d := int64(0)
		// count x distance
		for i := min(combo[0].X, combo[1].X); i < max(combo[0].X, combo[1].X); i++ {
			if galaxyColumns[i] {
				d++
			} else {
				d += factor
			}
		}

		// count y distance
		for j := min(combo[0].Y, combo[1].Y); j < max(combo[0].Y, combo[1].Y); j++ {
			if galaxyRows[j] {
				d++
			} else {
				d += factor
			}
		}

		sum += d
	}
	return sum
}
