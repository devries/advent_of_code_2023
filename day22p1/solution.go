package day22p1

import (
	"fmt"
	"io"
	"strconv"
	"strings"

	"aoc/utils"
)

func Solve(r io.Reader) any {
	lines := utils.ReadLines(r)
	var bricks []Brick

	for _, ln := range lines {
		b := parseLine(ln)
		bricks = append(bricks, b)
	}

	for _, b := range bricks {
		fmt.Println(b)
	}

	return 0
}

type Vector struct {
	X int
	Y int
	Z int
}

type Brick struct {
	A Vector
	B Vector
}

func parseLine(ln string) Brick {
	parts := strings.Split(ln, "~")

	return Brick{parseVector(parts[0]), parseVector(parts[1])}

}

func parseVector(v string) Vector {
	parts := strings.Split(v, ",")

	vec := Vector{}

	for i, s := range parts {
		val, err := strconv.Atoi(s)
		if err != nil {
			utils.Check(err, "Unable to convert %s into integer", s)
		}
		switch i {
		case 0:
			vec.X = val
		case 1:
			vec.Y = val
		case 2:
			vec.Z = val
		}
	}

	return vec
}
