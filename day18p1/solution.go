package day18p1

import (
	"fmt"
	"io"
	"regexp"
	"strconv"

	"aoc/utils"
)

func Solve(r io.Reader) any {
	lines := utils.ReadLines(r)

	re := regexp.MustCompile(`([RDLU])\s+(\d+)\s+\(#([0-9a-f]+)\)`)

	grid := make(map[utils.Point]bool)
	pos := utils.Point{}
	grid[pos] = true
	xmin, ymin, xmax, ymax := 0, 0, 0, 0

	for _, ln := range lines {
		sm := re.FindStringSubmatch(ln)

		inst := Instruction{}
		switch sm[1] {
		case "U":
			inst.Direction = utils.North
		case "D":
			inst.Direction = utils.South
		case "R":
			inst.Direction = utils.East
		case "L":
			inst.Direction = utils.West
		default:
			panic("Direction not good")
		}

		var err error
		inst.Distance, err = strconv.Atoi(sm[2])
		if err != nil {
			utils.Check(err, "Unable to convert %s to int", sm[2])
		}

		inst.Color = sm[3]

		// dig out trenches
		for i := 0; i < inst.Distance; i++ {
			pos = pos.Add(inst.Direction)
			if pos.X > xmax {
				xmax = pos.X
			}
			if pos.Y > ymax {
				ymax = pos.Y
			}
			if pos.X < xmin {
				xmin = pos.X
			}
			if pos.Y < ymin {
				ymin = pos.Y
			}

			grid[pos] = true
		}
	}

	if utils.Verbose {
		printGrid(grid, xmin, ymin, xmax, ymax)
	}

	g := &Grid{grid, xmin, ymin, xmax, ymax}

	bfs := utils.NewBFS[utils.Point]()

	_, err := bfs.Run(g)
	if err != utils.BFSNotFound {
		panic("did not exhaust grid")
	}

	outside := len(bfs.Visited)

	if utils.Verbose {
		printGrid(bfs.Visited, xmin-1, ymin-1, xmax+1, ymax+1)
	}

	area := (xmax - xmin + 3) * (ymax - ymin + 3)
	return area - outside
}

type Instruction struct {
	Direction utils.Point
	Distance  int
	Color     string
}

func printGrid(grid map[utils.Point]bool, xmin, ymin, xmax, ymax int) {
	for j := ymax; j >= ymin; j-- {
		for i := xmin; i <= xmax; i++ {
			switch grid[utils.Point{X: i, Y: j}] {
			case true:
				fmt.Printf("#")
			case false:
				fmt.Printf(".")
			}
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}

type Grid struct {
	Edges map[utils.Point]bool
	Xmin  int
	Ymin  int
	Xmax  int
	Ymax  int
}

func (g *Grid) GetInitial() utils.Point {
	return utils.Point{X: g.Xmin - 1, Y: g.Ymin - 1}
}

func (g *Grid) GetNeighbors(p utils.Point) []utils.Point {
	ret := []utils.Point{}

	for _, dir := range utils.Directions {
		np := p.Add(dir)

		if np.X > g.Xmax+1 || np.X < g.Xmin-1 || np.Y < g.Ymin-1 || np.Y > g.Ymax+1 {
			continue
		}
		if !g.Edges[np] {
			ret = append(ret, np)
		}
	}

	return ret
}

func (g *Grid) IsFinal(p utils.Point) bool {
	return false
}
