package day15p2

import (
	"io"
	"strconv"
	"strings"

	"aoc/utils"
)

func Solve(r io.Reader) any {
	lines := utils.ReadLines(r)

	steps := strings.Split(lines[0], ",")

	boxes := make(map[byte][]Lens)
	for _, s := range steps {
		if strings.Contains(s, "=") {
			parts := strings.Split(s, "=")
			fl, err := strconv.Atoi(parts[1])
			utils.Check(err, "unable to convert %s to integer", parts[1])

			h := elfHash(parts[0])

			lenses := boxes[h]
			found := false
			for i, l := range lenses {
				if l.Label == parts[0] {
					found = true
					l.FocalLength = fl
					lenses[i] = l
					break
				}
			}

			if !found {
				newlenses := make([]Lens, len(lenses)+1)
				copy(newlenses, lenses)
				newlenses[len(lenses)] = Lens{parts[0], fl}
				boxes[h] = newlenses
			}
		} else {
			label, _ := strings.CutSuffix(s, "-")

			h := elfHash(label)

			lenses := boxes[h]

			newlenses := make([]Lens, 0, len(lenses))
			for _, l := range lenses {
				if l.Label != label {
					newlenses = append(newlenses, l)
				}
			}

			boxes[h] = newlenses
		}
	}

	// Find focusing power
	var sum uint64

	for i := 0; i < 256; i++ {
		lenses := boxes[byte(i)]
		if len(lenses) != 0 {
			for j, l := range lenses {
				boxpower := (uint64(i) + 1) * uint64((j+1)*l.FocalLength)
				sum += boxpower
			}
		}
	}
	return sum
}

func elfHash(s string) byte {
	var res byte

	for _, r := range s {
		res += byte(r)
		res *= 17
	}

	return res
}

type Lens struct {
	Label       string
	FocalLength int
}
