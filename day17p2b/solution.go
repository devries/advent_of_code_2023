package day17p2b

import (
	"io"

	"aoc/utils"
)

func Solve(r io.Reader) any {
	lines := utils.ReadLines(r)

	m := NewMaze(lines)

	dij := utils.NewDijkstra[State]()

	finalState, err := dij.Run(m)
	utils.Check(err, "error in search")

	return dij.Distance[finalState]
}

type Maze struct {
	HeatLoss    map[utils.Point]uint64
	BottomRight utils.Point
}

func NewMaze(lines []string) *Maze {
	m := Maze{HeatLoss: make(map[utils.Point]uint64)}

	for j, ln := range lines {
		for i, r := range ln {
			p := utils.Point{X: i, Y: -j}
			v := uint64(r - '0')
			m.HeatLoss[p] = v
			m.BottomRight = p
		}
	}

	return &m
}

type State struct {
	Position  utils.Point
	Direction utils.Point
}

var zero = utils.Point{X: 0, Y: 0}

func (m *Maze) GetInitial() State {
	// special starting direction from which we will go East to South
	return State{zero, zero}
}

func (m *Maze) GetEdges(s State) []utils.Edge[State] {
	ret := []utils.Edge[State]{}

	nextDirections := make([]utils.Point, 2)

	switch s.Direction {
	case zero:
		// This is the special starting state
		nextDirections[0] = utils.East
		nextDirections[1] = utils.South
	case utils.East, utils.West:
		nextDirections[0] = utils.South
		nextDirections[1] = utils.North
	case utils.North, utils.South:
		nextDirections[0] = utils.East
		nextDirections[1] = utils.West
	}

	// Left and Right
	for _, dir := range nextDirections {
		p := s.Position
		var length uint64
		for step := 1; step <= 10; step++ {
			p = p.Add(dir)
			if delta, ok := m.HeatLoss[p]; ok {
				length += delta
				if step >= 4 {
					ret = append(ret, utils.Edge[State]{Node: State{p, dir}, Distance: length})
				}
			}
		}
	}

	return ret
}

func (m *Maze) IsFinal(s State) bool {
	if s.Position == m.BottomRight {
		return true
	}

	return false
}
