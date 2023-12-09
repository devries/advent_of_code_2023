package day09p2

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
		parts := strings.Fields(ln)

		values := []int{}

		for _, p := range parts {
			v, err := strconv.Atoi(p)
			utils.Check(err, "Unable to convert %s to int", p)

			values = append(values, v)
		}
		// Start differentiating
		valueRows := [][]int{values}
		for {
			newvalues := make([]int, 0, len(values)-1)

			allzeros := true
			for i := 1; i < len(values); i++ {
				nv := values[i] - values[i-1]
				if nv != 0 {
					allzeros = false
				}
				newvalues = append(newvalues, nv)
			}
			valueRows = append(valueRows, newvalues)
			values = newvalues

			if allzeros {
				break
			}
		}

		delta := 0
		for i := len(valueRows) - 1; i >= 0; i-- {
			delta = valueRows[i][0] - delta
		}
		sum += delta
	}

	return sum
}
