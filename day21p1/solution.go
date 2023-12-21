package day21p1

import (
	"io"

	"aoc/utils"
)

func Solve(r io.Reader) any {
	return solveSteps(r, 64)
}

func solveSteps(r io.Reader, steps int) int {
	lines := utils.ReadLines(r)

	garden := Maze{make(map[utils.Point]rune), utils.NewBFS[utils.Point](), steps}

	for j, ln := range lines {
		for i, r := range ln {
			p := utils.Point{X: i, Y: -j}
			garden.positions[p] = r
		}
	}

	_, err := garden.bfs.Run(garden)
	if err != utils.BFSNotFound {
		panic("this should end when points are exhausted")
	}

	count := 0
	for _, d := range garden.bfs.Distance {
		if d%2 == 0 {
			count++
		}
	}

	return count
}

type Maze struct {
	positions map[utils.Point]rune
	bfs       *utils.BreadthFirstSearch[utils.Point]
	steps     int
}

func (m Maze) GetInitial() utils.Point {
	for k, v := range m.positions {
		if v == 'S' {
			return k
		}
	}
	panic("unable to find starting point")
}

func (m Maze) GetNeighbors(pos utils.Point) []utils.Point {
	ret := []utils.Point{}

	if m.bfs.Distance[pos] < uint64(m.steps) {
		for _, dir := range utils.Directions {
			np := pos.Add(dir)
			if m.positions[np] == '.' {
				ret = append(ret, np)
			}
		}
	}

	return ret
}

func (m Maze) IsFinal(pos utils.Point) bool {
	return false
}
